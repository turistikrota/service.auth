package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.auth/domains/user"
)

type CheckEmailQuery struct {
	Email string `json:"email" validate:"required,email"`
}

type CheckEmailRes struct {
	Exists bool `json:"exists"`
}

type CheckEmailHandler cqrs.HandlerFunc[CheckEmailQuery, *CheckEmailRes]

func NewCheckEmailHandler(repo user.Repo) CheckEmailHandler {
	return func(ctx context.Context, query CheckEmailQuery) (*CheckEmailRes, *i18np.Error) {
		res, err := repo.CheckEmail(ctx, query.Email)
		if err != nil {
			return nil, err
		}
		return &CheckEmailRes{
			Exists: res,
		}, nil
	}
}
