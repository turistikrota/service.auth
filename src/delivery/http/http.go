package http

import (
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/server/http"
	"github.com/mixarchitecture/microp/server/http/parser"
	"github.com/mixarchitecture/microp/validator"
	"github.com/ssibrahimbas/turnstile"
	"github.com/turistikrota/service.auth/src/app"
	"github.com/turistikrota/service.auth/src/config"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/csrf"
	"github.com/turistikrota/service.shared/server/http/auth"
	"github.com/turistikrota/service.shared/server/http/auth/claim_guard"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.shared/server/http/auth/device_uuid"
	"github.com/turistikrota/service.shared/server/http/auth/refresh_token"
	"github.com/turistikrota/service.shared/server/http/auth/required_access"
	turnstile_middleware "github.com/turistikrota/service.shared/server/http/auth/turnstile"
	"github.com/turistikrota/service.shared/server/http/auth/two_factor"
)

type Server struct {
	app          app.Application
	i18n         i18np.I18n
	validator    validator.Validator
	ctx          context.Context
	tknSrv       token.Service
	sessionSrv   session.Service
	config       config.App
	turnstileSrv turnstile.Service
}

type Config struct {
	App          app.Application
	I18n         i18np.I18n
	Validator    validator.Validator
	Context      context.Context
	Config       config.App
	TokenSrv     token.Service
	SessionSrv   session.Service
	TurnstileSrv turnstile.Service
}

func New(config Config) Server {
	return Server{
		app:          config.App,
		i18n:         config.I18n,
		validator:    config.Validator,
		ctx:          config.Context,
		tknSrv:       config.TokenSrv,
		sessionSrv:   config.SessionSrv,
		config:       config.Config,
		turnstileSrv: config.TurnstileSrv,
	}
}

func (h Server) Load(router fiber.Router) fiber.Router {
	router.Use(h.cors(), h.deviceUUID())
	router.Post("/register", h.rateLimit(10), h.turnstile(), h.wrapWithTimeout(h.Register))
	router.Post("/login", h.rateLimit(10), h.turnstile(), h.wrapWithTimeout(h.Login))
	router.Post("/checkEmail", h.rateLimit(10), h.turnstile(), h.wrapWithTimeout(h.CheckEmail))
	router.Post("/logout", h.rateLimit(10), h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.Logout))
	router.Put("/refresh", h.rateLimit(20), h.currentUserRefresh(), h.requiredRefreshToken(), h.wrapWithTimeout(h.RefreshToken))
	router.Get("/2fa/check", h.rateLimit(10), h.currentUserTemp(), h.twoFactor(), h.wrapWithTimeout(h.LoginVerified))
	router.Post("/re-verify", h.rateLimit(10), h.turnstile(), h.wrapWithTimeout(h.ReSendVerification))
	router.Post("/:token", h.rateLimit(10), h.turnstile(), h.wrapWithTimeout(h.Verify))
	router.Get("/", h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.CurrentUser))
	router.Delete("/", h.currentUserAccess(), h.requiredAccess(), h.turnstile(), h.wrapWithTimeout(h.UserDelete))
	router.Get("/user-list", h.currentUserAccess(), h.requiredAccess(), h.adminRoute(config.Roles.UserList), h.wrapWithTimeout(h.UserList))

	session := router.Group("/session", h.currentUserAccess(), h.requiredAccess())
	session.Get("/", h.wrapWithTimeout(h.SessionList))
	session.Delete("/others", h.wrapWithTimeout(h.SessionDestroyOthers))
	session.Delete("/all", h.wrapWithTimeout(h.SessionDestroyAll))
	session.Delete("/:device_uuid", h.wrapWithTimeout(h.SessionDestroy))

	return router
}

func (h Server) parseBody(c *fiber.Ctx, d interface{}) {
	parser.ParseBody(c, h.validator, h.i18n, d)
}

func (h Server) parseParams(c *fiber.Ctx, d interface{}) {
	parser.ParseParams(c, h.validator, h.i18n, d)
}

func (h Server) parseQuery(c *fiber.Ctx, d interface{}) {
	parser.ParseQuery(c, h.validator, h.i18n, d)
}

func (h Server) wrapWithTimeout(fn fiber.Handler) fiber.Handler {
	return timeout.NewWithContext(fn, 10*time.Second)
}

func (h Server) rateLimit(limit int) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        limit,
		Expiration: 3 * time.Minute,
	})
}

func (h Server) removeSelectedAccountInCookie(ctx *fiber.Ctx) {
	ctx.Cookie(&fiber.Cookie{
		Name:     ".s.a.u",
		Value:    "",
		Expires:  time.Now().Add(time.Hour * -1),
		HTTPOnly: true,
		Secure:   true,
		Domain:   h.config.HttpHeaders.Domain,
		SameSite: "Strict",
	})
}

func (h Server) csrf() fiber.Handler {
	return csrf.New(csrf.Config{
		Base: csrf.Base{
			SameSite:   h.config.CSRF.BaseEnv.SameSite,
			Domain:     h.config.CSRF.BaseEnv.Domain,
			Secure:     h.config.CSRF.BaseEnv.Secure,
			HttpOnly:   h.config.CSRF.BaseEnv.HttpOnly,
			Expiration: h.config.CSRF.BaseEnv.Expiration,
		},
		Redis: csrf.EnvRedis{
			Host: h.config.CSRF.Redis.Host,
			Port: h.config.CSRF.Redis.Port,
			Pw:   h.config.CSRF.Redis.Pw,
			Db:   h.config.CSRF.Redis.Db,
		},
	})
}

func (h Server) turnstile() fiber.Handler {
	return turnstile_middleware.New(turnstile_middleware.Config{
		Service:            h.turnstileSrv,
		I18n:               h.i18n,
		BadRequestMsgKey:   Messages.Error.TurnstileBadRequest,
		UnauthorizedMsgKey: Messages.Error.TurnstileUnauthorized,
		Skip:               h.config.Turnstile.Skip,
		CheckForMobile:     true,
	})
}

func (h Server) currentUserRefresh() fiber.Handler {
	return current_user.New(current_user.Config{
		TokenSrv:   h.tknSrv,
		SessionSrv: h.sessionSrv,
		I18n:       &h.i18n,
		MsgKey:     "errors_auth_current_user",
		HeaderKey:  http.Headers.RefreshToken,
		CookieKey:  auth.Cookies.RefreshToken,
		UseCookie:  true,
		UseBearer:  true,
		IsRefresh:  true,
		IsAccess:   false,
		LocalKey:   refresh_token.LocalKey,
	})
}

func (h Server) currentUserTemp() fiber.Handler {
	return current_user.New(current_user.Config{
		TokenSrv:   h.tknSrv,
		SessionSrv: h.sessionSrv,
		I18n:       &h.i18n,
		MsgKey:     "errors_auth_current_user",
		HeaderKey:  http.Headers.TwoFactorToken,
		CookieKey:  auth.Cookies.TwoFactorToken,
		UseCookie:  true,
		UseBearer:  true,
		IsRefresh:  false,
		Is2FA:      true,
		IsAccess:   false,
		LocalKey:   two_factor.LocalKey,
	})
}

func (h Server) adminRoute(extra ...string) fiber.Handler {
	claims := []string{config.Roles.Admin}
	if len(extra) > 0 {
		claims = append(claims, extra...)
	}
	return claim_guard.New(claim_guard.Config{
		Claims: claims,
		I18n:   h.i18n,
		MsgKey: Messages.Error.AdminRoute,
	})
}

func (h Server) currentUserAccess() fiber.Handler {
	return current_user.New(current_user.Config{
		TokenSrv:   h.tknSrv,
		SessionSrv: h.sessionSrv,
		I18n:       &h.i18n,
		MsgKey:     "errors_auth_current_user",
		HeaderKey:  http.Headers.Authorization,
		CookieKey:  auth.Cookies.AccessToken,
		UseCookie:  true,
		UseBearer:  true,
		IsRefresh:  false,
		IsAccess:   true,
	})
}

func (h Server) twoFactor() fiber.Handler {
	return two_factor.New(two_factor.Config{
		I18n:   h.i18n,
		MsgKey: "errors_auth_two_factor",
	})
}

func (h Server) requiredAccess() fiber.Handler {
	return required_access.New(required_access.Config{
		I18n:   h.i18n,
		MsgKey: "errors_auth_required_auth",
	})
}

func (h Server) deviceUUID() fiber.Handler {
	return device_uuid.New(device_uuid.Config{
		Domain: h.config.HttpHeaders.Domain,
	})
}

func (h Server) requiredRefreshToken() fiber.Handler {
	return refresh_token.New(refresh_token.Config{
		I18n:   h.i18n,
		MsgKey: "errors_auth_refresh_token",
	})
}

func (h Server) cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     h.config.HttpHeaders.AllowedOrigins,
		AllowMethods:     h.config.HttpHeaders.AllowedMethods,
		AllowHeaders:     h.config.HttpHeaders.AllowedHeaders,
		AllowCredentials: h.config.HttpHeaders.AllowCredentials,
		AllowOriginsFunc: func(origin string) bool {
			origins := strings.Split(h.config.HttpHeaders.AllowedOrigins, ",")
			for _, o := range origins {
				if strings.Contains(origin, o) {
					return true
				}
			}
			return false
		},
	})
}
