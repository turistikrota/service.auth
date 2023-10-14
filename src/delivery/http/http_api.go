package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mileusna/useragent"
	httpI18n "github.com/mixarchitecture/microp/server/http/i18n"
	"github.com/mixarchitecture/microp/server/http/result"
	"github.com/turistikrota/service.auth/src/app/command"
	"github.com/turistikrota/service.auth/src/app/query"
	"github.com/turistikrota/service.auth/src/delivery/http/dto"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/server/http/auth"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.shared/server/http/auth/device_uuid"
	"github.com/turistikrota/service.shared/server/http/auth/refresh_token"
	"github.com/turistikrota/service.shared/server/http/auth/two_factor"
)

func (h Server) Register(ctx *fiber.Ctx) error {
	d := dto.Request.Register()
	h.parseBody(ctx, d)
	l, a := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	_, err := h.app.Commands.Register.Handle(ctx.UserContext(), d.ToCommand(l))
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.Success(Messages.Success.Register)
}

func (h Server) LoginVerified(ctx *fiber.Ctx) error {
	d := dto.Request.LoginVerified()
	l, a := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	d.UserUUID = current_user.Parse(ctx).UUID
	d.Device = h.makeDevice(ctx)
	d.DeviceUUID = device_uuid.Parse(ctx)
	res, err := h.app.Commands.LoginVerified.Handle(ctx.UserContext(), d.ToCommand())
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	refresh_token.Set(ctx, h.config.HttpHeaders.Domain, res.RefreshToken)
	current_user.SetCookie(ctx, auth.Cookies.AccessToken, h.config.HttpHeaders.Domain, res.AccessToken)
	return result.SuccessDetail(Messages.Success.Login, dto.Response.LoggedIn(res.AccessToken))
}

func (h Server) CheckEmail(ctx *fiber.Ctx) error {
	d := dto.Request.CheckEmail()
	h.parseBody(ctx, d)
	res, err := h.app.Queries.CheckEmail.Handle(ctx.UserContext(), d.ToQuery())
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.EmailAvailable, func() interface{} {
		return dto.Response.CheckEmail(res)
	})
}

func (h Server) Login(ctx *fiber.Ctx) error {
	d := dto.Request.Login()
	h.parseBody(ctx, d)
	l, a := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	d.Device = h.makeDevice(ctx)
	d.DeviceUUID = device_uuid.Parse(ctx)
	res, err := h.app.Commands.Login.Handle(ctx.UserContext(), d.ToCommand())
	if err != nil {
		if err.IsDetails() {
			return result.ErrorDetail(h.i18n.TranslateFromError(*err, l, a), err.Details)
		}
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	if res.Verify {
		two_factor.Set(ctx, h.config.HttpHeaders.Domain, res.TempToken)
		return result.ErrorDetail(Messages.Error.LoginVerify, dto.Response.VerifyRequired())
	}
	refresh_token.Set(ctx, h.config.HttpHeaders.Domain, res.RefreshToken)
	current_user.SetCookie(ctx, auth.Cookies.AccessToken, h.config.HttpHeaders.Domain, res.AccessToken)
	return result.SuccessDetail(Messages.Success.Login, dto.Response.LoggedIn(res.AccessToken))
}

func (h Server) Logout(ctx *fiber.Ctx) error {
	d := dto.Request.Logout()
	u := current_user.Parse(ctx)
	d.UserUUID = u.UUID
	d.DeviceUUID = device_uuid.Parse(ctx)
	_, err := h.app.Commands.Logout.Handle(ctx.UserContext(), d.ToCommand())
	l, a := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	refresh_token.Remove(ctx, h.config.HttpHeaders.Domain)
	current_user.RemoveCookie(ctx, auth.Cookies.AccessToken, h.config.HttpHeaders.Domain)
	h.removeSelectedAccountInCookie(ctx)
	return result.Success(Messages.Success.Logout)
}

func (h Server) UserDelete(ctx *fiber.Ctx) error {
	userId := current_user.Parse(ctx).UUID
	_, err := h.app.Commands.UserDelete.Handle(ctx.UserContext(), command.UserDeleteCommand{
		UserID: userId,
	})
	l, a := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	refresh_token.Remove(ctx, h.config.HttpHeaders.Domain)
	current_user.RemoveCookie(ctx, auth.Cookies.AccessToken, h.config.HttpHeaders.Domain)
	h.removeSelectedAccountInCookie(ctx)
	return result.Success(Messages.Success.Logout)
}

func (h Server) RefreshToken(ctx *fiber.Ctx) error {
	d := dto.Request.RefreshToken()
	d.UserUUID = current_user.Parse(ctx).UUID
	d.DeviceUUID = device_uuid.Parse(ctx)
	d.IpAddress = ctx.IP()
	d.RefreshToken = refresh_token.Parse(ctx)
	d.AccessToken = current_user.GetAccessTokenFromCookie(ctx)
	l, a := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	res, err := h.app.Commands.RefreshToken.Handle(ctx.UserContext(), d.ToCommand())
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	refresh_token.Set(ctx, h.config.HttpHeaders.Domain, res.RefreshToken)
	current_user.SetCookie(ctx, auth.Cookies.AccessToken, h.config.HttpHeaders.Domain, res.AccessToken)
	return result.SuccessDetail(Messages.Success.Extend, dto.Response.LoggedIn(res.AccessToken))
}

func (h Server) Verify(ctx *fiber.Ctx) error {
	d := dto.Request.Verify()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.Verify.Handle(ctx.UserContext(), d.ToCommand())
	return result.IfSuccessParams(err, ctx, h.i18n, Messages.Success.Verify)
}

func (h Server) ReSendVerification(ctx *fiber.Ctx) error {
	d := dto.Request.ReSendVerification()
	h.parseBody(ctx, d)
	l, _ := httpI18n.GetLanguagesInContext(h.i18n, ctx)
	_, err := h.app.Commands.ReSendVerification.Handle(ctx.UserContext(), d.ToCommand(l))
	return result.IfSuccess(err, ctx, h.i18n, Messages.Success.ReSendVerification)
}

func (h Server) SessionDestroy(ctx *fiber.Ctx) error {
	d := dto.Request.Device()
	h.parseParams(ctx, d)
	_, err := h.app.Commands.SessionDestroy.Handle(ctx.UserContext(), d.ToDestroyCommand(current_user.Parse(ctx).UUID))
	return result.IfSuccessParams(err, ctx, h.i18n, Messages.Success.SessionDestroy)
}

func (h Server) SessionDestroyOthers(ctx *fiber.Ctx) error {
	_, err := h.app.Commands.SessionDestroyOthers.Handle(ctx.UserContext(), command.SessionDestroyOthersCommand{
		UserUUID:   current_user.Parse(ctx).UUID,
		DeviceUUID: device_uuid.Parse(ctx),
	})
	return result.IfSuccessParams(err, ctx, h.i18n, Messages.Success.SessionDestroyOthers)
}

func (h Server) SessionList(ctx *fiber.Ctx) error {
	res, err := h.app.Queries.SessionList.Handle(ctx.UserContext(), query.SessionListQuery{
		UserUUID: current_user.Parse(ctx).UUID,
	})
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.SessionList, func() interface{} {
		return dto.Response.SessionList(res)
	})
}

func (h Server) SessionDestroyAll(ctx *fiber.Ctx) error {
	_, err := h.app.Commands.SessionDestroyAll.Handle(ctx.UserContext(), command.SessionDestroyAllCommand{
		UserUUID: current_user.Parse(ctx).UUID,
	})
	return result.IfSuccessParams(err, ctx, h.i18n, Messages.Success.SessionDestroyAll)
}

func (h Server) CurrentUser(ctx *fiber.Ctx) error {
	u := current_user.Parse(ctx)
	res := dto.Response.CurrentUser(u)
	return result.SuccessDetail(Messages.Success.CurrentUser, res)
}

func (h Server) UserList(ctx *fiber.Ctx) error {
	d := dto.Request.Pagination()
	h.parseQuery(ctx, d)
	d.Default()
	res, err := h.app.Queries.UserList.Handle(ctx.UserContext(), query.UserListQuery{
		Offset: int64(*d.Page-1) * *d.Limit,
		Limit:  *d.Limit,
	})
	return result.IfSuccessDetail(err, ctx, h.i18n, Messages.Success.UserList, func() interface{} {
		return dto.Response.UserList(res)
	})
}

func (h Server) makeDevice(ctx *fiber.Ctx) *session.Device {
	ua := useragent.Parse(ctx.Get("User-Agent"))
	t := "desktop"
	if ua.Mobile {
		t = "mobile"
	} else if ua.Tablet {
		t = "tablet"
	} else if ua.Bot {
		t = "bot"
	}
	ip := ctx.Get("CF-Connecting-IP")
	if ip == "" {
		ip = ctx.IP()
	}
	return &session.Device{
		Name: ua.Name,
		Type: t,
		OS:   ua.OS,
		IP:   ip,
	}
}
