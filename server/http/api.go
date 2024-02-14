package http

import (
	"fmt"

	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/mileusna/useragent"
	"github.com/turistikrota/service.auth/app/command"
	"github.com/turistikrota/service.auth/app/query"
	"github.com/turistikrota/service.auth/domains/user"
	"github.com/turistikrota/service.auth/pkg/utils"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/server/http/auth"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
	"github.com/turistikrota/service.shared/server/http/auth/device_uuid"
	"github.com/turistikrota/service.shared/server/http/auth/refresh_token"
)

func (h srv) Register(ctx *fiber.Ctx) error {
	cmd := command.RegisterCmd{}
	h.parseBody(ctx, &cmd)
	l, a := i18n.ParseLocales(ctx)
	cmd.Lang = l
	res, err := h.app.Commands.Register(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) CheckEmail(ctx *fiber.Ctx) error {
	fmt.Println("CheckEmail")
	query := query.CheckEmailQuery{}
	h.parseBody(ctx, &query)
	res, err := h.app.Queries.CheckEmail(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) Login(ctx *fiber.Ctx) error {
	cmd := command.LoginCmd{}
	h.parseBody(ctx, &cmd)
	l, a := i18n.ParseLocales(ctx)
	cmd.Device = h.makeDevice(ctx)
	cmd.DeviceUUID = device_uuid.Parse(ctx)
	res, err := h.app.Commands.Login(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	//refresh_token.Set(ctx, h.config.HttpHeaders.Domain, res.RefreshToken)
	current_user.SetCookie(ctx, auth.Cookies.RefreshToken, h.config.HttpHeaders.Domain, res.RefreshToken)
	current_user.SetCookie(ctx, auth.Cookies.AccessToken, h.config.HttpHeaders.Domain, res.AccessToken)
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) Logout(ctx *fiber.Ctx) error {
	cmd := command.LogoutCmd{}
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.DeviceUUID = device_uuid.Parse(ctx)
	res, err := h.app.Commands.Logout(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	//refresh_token.Remove(ctx, h.config.HttpHeaders.Domain)
	current_user.RemoveCookie(ctx, auth.Cookies.RefreshToken, h.config.HttpHeaders.Domain)
	current_user.RemoveCookie(ctx, auth.Cookies.AccessToken, h.config.HttpHeaders.Domain)
	h.removeSelectedAccountInCookie(ctx)
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) UserDelete(ctx *fiber.Ctx) error {
	cmd := command.UserDeleteCmd{}
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.DeviceUUID = device_uuid.Parse(ctx)
	res, err := h.app.Commands.UserDelete(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	//refresh_token.Remove(ctx, h.config.HttpHeaders.Domain)
	current_user.RemoveCookie(ctx, auth.Cookies.RefreshToken, h.config.HttpHeaders.Domain)
	current_user.RemoveCookie(ctx, auth.Cookies.AccessToken, h.config.HttpHeaders.Domain)
	h.removeSelectedAccountInCookie(ctx)
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) RefreshToken(ctx *fiber.Ctx) error {
	cmd := command.RefreshTokenCmd{}
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.DeviceUUID = device_uuid.Parse(ctx)
	cmd.IpAddress = ctx.IP()
	cmd.RefreshToken = refresh_token.Parse(ctx)
	cmd.AccessToken = current_user.GetAccessTokenFromCookie(ctx)
	res, err := h.app.Commands.RefreshToken(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	//refresh_token.Set(ctx, h.config.HttpHeaders.Domain, res.RefreshToken)
	current_user.SetCookie(ctx, auth.Cookies.RefreshToken, h.config.HttpHeaders.Domain, res.RefreshToken)
	current_user.SetCookie(ctx, auth.Cookies.AccessToken, h.config.HttpHeaders.Domain, res.AccessToken)
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ReSendVerificationCode(ctx *fiber.Ctx) error {
	cmd := command.ReSendVerificationCodeCmd{}
	h.parseBody(ctx, &cmd)
	l, a := i18n.ParseLocales(ctx)
	cmd.Lang = l
	res, err := h.app.Commands.ReSendVerificationCode(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) SessionDestroy(ctx *fiber.Ctx) error {
	cmd := command.SessionDestroyCmd{}
	h.parseParams(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.SessionDestroy(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) SessionDestroyAll(ctx *fiber.Ctx) error {
	cmd := command.SessionDestroyAllCmd{}
	cmd.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.SessionDestroyAll(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) Verify(ctx *fiber.Ctx) error {
	cmd := command.VerifyCmd{}
	h.parseParams(ctx, &cmd)
	res, err := h.app.Commands.Verify(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) SessionDestroyOthers(ctx *fiber.Ctx) error {
	cmd := command.SessionDestroyOthersCmd{}
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.DeviceUUID = device_uuid.Parse(ctx)
	res, err := h.app.Commands.SessionDestroyOthers(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) GetCurrentUser(ctx *fiber.Ctx) error {
	u := current_user.Parse(ctx)
	return result.SuccessDetail(Messages.Success.Ok, map[string]interface{}{
		"uuid":       u.UUID,
		"email":      u.Email,
		"roles":      u.Roles,
		"accounts":   u.Accounts,
		"businesses": u.Businesses,
	})
}

func (h srv) UserList(ctx *fiber.Ctx) error {
	pagi := utils.Pagination{}
	h.parseQuery(ctx, &pagi)
	filters := user.FilterEntity{}
	h.parseQuery(ctx, &filters)
	query := query.UserListQuery{
		Pagination:   &pagi,
		FilterEntity: &filters,
	}
	res, err := h.app.Queries.UserList(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) SessionList(ctx *fiber.Ctx) error {
	query := query.SessionListQuery{
		UserUUID:   current_user.Parse(ctx).UUID,
		DeviceUUID: device_uuid.Parse(ctx),
	}
	res, err := h.app.Queries.SessionList(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Sessions)
}

func (h srv) SetFcmToken(ctx *fiber.Ctx) error {
	cmd := command.SetFcmTokenCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.DeviceUUID = device_uuid.Parse(ctx)
	res, err := h.app.Commands.SetFcmToken(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) ChangePassword(ctx *fiber.Ctx) error {
	cmd := command.ChangePasswordCmd{}
	h.parseBody(ctx, &cmd)
	cmd.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.ChangePassword(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
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
