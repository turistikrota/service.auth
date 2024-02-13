package http

import (
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.auth/app/command"
	"github.com/turistikrota/service.auth/app/query"
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
	cmd := query.CheckEmailQuery{}
	h.parseBody(ctx, &cmd)
	res, err := h.app.Queries.CheckEmail(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(*h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}
