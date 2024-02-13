package http

import (
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/mileusna/useragent"
	"github.com/turistikrota/service.auth/app/command"
	"github.com/turistikrota/service.auth/app/query"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/server/http/auth"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.shared/server/http/auth/device_uuid"
	"github.com/turistikrota/service.shared/server/http/auth/refresh_token"
)

func (h srv) Register(ctx *fiber.Ctx) error {
	cmd := command.RegisterCmd{}
	h.parseBody(ctx, &cmd)
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	cmd.Lang = l
	res, err := h.app.Commands.Register(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) CheckEmail(ctx *fiber.Ctx) error {
	query := query.CheckEmailQuery{}
	h.parseBody(ctx, &query)
	res, err := h.app.Queries.CheckEmail(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) Login(ctx *fiber.Ctx) error {
	cmd := command.LoginCmd{}
	h.parseBody(ctx, &cmd)
	l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
	cmd.Device = h.makeDevice(ctx)
	cmd.DeviceUUID = device_uuid.Parse(ctx)
	res, err := h.app.Commands.Login(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	refresh_token.Set(ctx, h.config.HttpHeaders.Domain, res.RefreshToken)
	current_user.SetCookie(ctx, auth.Cookies.AccessToken, h.config.HttpHeaders.Domain, res.AccessToken)
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) Logout(ctx *fiber.Ctx) error {
	cmd := command.LogoutCmd{}
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.DeviceUUID = device_uuid.Parse(ctx)
	res, err := h.app.Commands.Logout(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	refresh_token.Remove(ctx, h.config.HttpHeaders.Domain)
	current_user.RemoveCookie(ctx, auth.Cookies.AccessToken, h.config.HttpHeaders.Domain)
	h.removeSelectedAccountInCookie(ctx)
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) makeDevice(ctx *fiber.Ctx) *session.Device {
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
