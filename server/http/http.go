package http

import (
	"fmt"
	"strings"
	"time"

	"github.com/cilloparch/cillop/helpers/http"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/server"
	"github.com/cilloparch/cillop/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/ssibrahimbas/turnstile"

	"github.com/turistikrota/service.auth/app"
	"github.com/turistikrota/service.auth/config"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	httpServer "github.com/turistikrota/service.shared/server/http"
	"github.com/turistikrota/service.shared/server/http/auth"
	"github.com/turistikrota/service.shared/server/http/auth/claim_guard"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.shared/server/http/auth/device_uuid"
	"github.com/turistikrota/service.shared/server/http/auth/refresh_token"
	"github.com/turistikrota/service.shared/server/http/auth/required_access"
	turnstile_middleware "github.com/turistikrota/service.shared/server/http/auth/turnstile"
)

type srv struct {
	config       config.App
	app          app.Application
	i18n         *i18np.I18n
	validator    validation.Validator
	tknSrv       token.Service
	sessionSrv   session.Service
	httpHeaders  config.HttpHeaders
	turnstileSrv turnstile.Service
}

type Config struct {
	Env          config.App
	App          app.Application
	I18n         *i18np.I18n
	Validator    validation.Validator
	HttpHeaders  config.HttpHeaders
	TokenSrv     token.Service
	SessionSrv   session.Service
	TurnstileSrv turnstile.Service
}

func New(config Config) server.Server {
	fmt.Println("http server", config)
	return srv{
		config:       config.Env,
		app:          config.App,
		i18n:         config.I18n,
		validator:    config.Validator,
		tknSrv:       config.TokenSrv,
		sessionSrv:   config.SessionSrv,
		httpHeaders:  config.HttpHeaders,
		turnstileSrv: config.TurnstileSrv,
	}
}

func (h srv) Listen() error {
	return http.RunServer(http.Config{
		Host:        h.config.Http.Host,
		Port:        h.config.Http.Port,
		I18n:        h.i18n,
		Debug:       true,
		AcceptLangs: []string{},
		CreateHandler: func(router fiber.Router) fiber.Router {
			router.Use(h.cors(), h.deviceUUID())

			router.Post("/register", h.rateLimit(10), h.turnstile(), h.wrapWithTimeout(h.Register))
			router.Post("/login", h.rateLimit(10), h.turnstile(), h.wrapWithTimeout(h.Login))
			router.Post("/checkEmail", h.logger(0), h.rateLimit(10), h.logger(1), h.turnstile(), h.logger(2), h.wrapWithTimeout(h.CheckEmail))
			router.Post("/logout", h.rateLimit(10), h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.Logout))
			router.Put("/refresh", h.rateLimit(20), h.currentUserRefresh(), h.requiredRefreshToken(), h.wrapWithTimeout(h.RefreshToken))
			router.Post("/re-verify", h.rateLimit(10), h.turnstile(), h.wrapWithTimeout(h.ReSendVerificationCode))
			router.Post("/:token", h.rateLimit(10), h.turnstile(), h.wrapWithTimeout(h.Verify))
			router.Get("/", h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.GetCurrentUser))
			router.Delete("/", h.currentUserAccess(), h.requiredAccess(), h.turnstile(), h.wrapWithTimeout(h.UserDelete))
			router.Get("/user-list", h.currentUserAccess(), h.requiredAccess(), h.adminRoute(config.Roles.User.List), h.wrapWithTimeout(h.UserList))
			router.Patch("/fcm", h.currentUserAccess(), h.requiredAccess(), h.wrapWithTimeout(h.SetFcmToken))
			router.Patch("/password", h.currentUserAccess(), h.requiredAccess(), h.turnstile(), h.wrapWithTimeout(h.ChangePassword))

			session := router.Group("/session", h.currentUserAccess(), h.requiredAccess())
			session.Get("/", h.wrapWithTimeout(h.SessionList))
			session.Delete("/others", h.wrapWithTimeout(h.SessionDestroyOthers))
			session.Delete("/all", h.wrapWithTimeout(h.SessionDestroyAll))
			session.Delete("/:device_uuid", h.wrapWithTimeout(h.SessionDestroy))
			return router
		},
	})
}

func (h srv) logger(log int) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Println("log", log)
		return ctx.Next()
	}
}

func (h srv) parseBody(c *fiber.Ctx, d interface{}) {
	http.ParseBody(c, h.validator, *h.i18n, d)
}

func (h srv) parseParams(c *fiber.Ctx, d interface{}) {
	http.ParseParams(c, h.validator, *h.i18n, d)
}

func (h srv) parseQuery(c *fiber.Ctx, d interface{}) {
	http.ParseQuery(c, h.validator, *h.i18n, d)
}

func (h srv) currentUserAccess() fiber.Handler {
	return current_user.New(current_user.Config{
		TokenSrv:   h.tknSrv,
		SessionSrv: h.sessionSrv,
		I18n:       h.i18n,
		MsgKey:     Messages.Error.CurrentUserAccess,
		HeaderKey:  httpServer.Headers.Authorization,
		CookieKey:  auth.Cookies.AccessToken,
		UseCookie:  true,
		UseBearer:  true,
		IsRefresh:  false,
		IsAccess:   true,
	})
}

func (h srv) rateLimit(limit int) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        limit,
		Expiration: 3 * time.Minute,
	})
}

func (h srv) deviceUUID() fiber.Handler {
	return device_uuid.New(device_uuid.Config{
		Domain: h.httpHeaders.Domain,
	})
}

func (h srv) requiredAccess() fiber.Handler {
	return required_access.New(required_access.Config{
		I18n:   *h.i18n,
		MsgKey: Messages.Error.RequiredAuth,
	})
}

func (h srv) adminRoute(extra ...string) fiber.Handler {
	claims := []string{config.Roles.Admin}
	if len(extra) > 0 {
		claims = append(claims, extra...)
	}
	return claim_guard.New(claim_guard.Config{
		Claims: claims,
		I18n:   *h.i18n,
		MsgKey: Messages.Error.AdminRoute,
	})
}

func (h srv) removeSelectedAccountInCookie(ctx *fiber.Ctx) {
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

func (h srv) turnstile() fiber.Handler {
	return turnstile_middleware.New(turnstile_middleware.Config{
		Service:        h.turnstileSrv,
		Skip:           h.config.Turnstile.Skip,
		CheckForMobile: true,
	})
}

func (h srv) currentUserRefresh() fiber.Handler {
	return current_user.New(current_user.Config{
		TokenSrv:   h.tknSrv,
		SessionSrv: h.sessionSrv,
		MsgKey:     "errors_auth_current_user",
		HeaderKey:  httpServer.Headers.RefreshToken,
		CookieKey:  auth.Cookies.RefreshToken,
		UseCookie:  true,
		UseBearer:  true,
		IsRefresh:  true,
		IsAccess:   false,
		LocalKey:   refresh_token.LocalKey,
	})
}

func (h srv) requiredRefreshToken() fiber.Handler {
	return refresh_token.New(refresh_token.Config{
		MsgKey: "errors_auth_refresh_token",
		I18n:   *h.i18n,
	})
}

func (h srv) cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowMethods:     h.httpHeaders.AllowedMethods,
		AllowHeaders:     h.httpHeaders.AllowedHeaders,
		AllowCredentials: h.httpHeaders.AllowCredentials,
		AllowOriginsFunc: func(origin string) bool {
			origins := strings.Split(h.httpHeaders.AllowedOrigins, ",")
			for _, o := range origins {
				if strings.Contains(origin, o) {
					return true
				}
			}
			return false
		},
	})
}

func (h srv) wrapWithTimeout(fn fiber.Handler) fiber.Handler {
	return timeout.NewWithContext(fn, 10*time.Second)
}
