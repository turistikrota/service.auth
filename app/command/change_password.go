package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.auth/domains/user"
	"github.com/turistikrota/service.shared/cipher"
)

type ChangePasswordCmd struct {
	OldPassword string `json:"old_password" validate:"required,password"`
	NewPassword string `json:"new_password" validate:"required,password"`
	UserUUID    string `json:"-"`
}

type ChangePasswordRes struct{}

type ChangePasswordHandler cqrs.HandlerFunc[ChangePasswordCmd, *ChangePasswordRes]

func NewChangePasswordHandler(repo user.Repo, factory user.Factory) ChangePasswordHandler {
	return func(ctx context.Context, cmd ChangePasswordCmd) (*ChangePasswordRes, *i18np.Error) {
		user, err := repo.GetByUUID(ctx, cmd.UserUUID)
		if err != nil {
			return nil, err
		}
		if err := cipher.Compare(user.Password, cmd.OldPassword); err != nil {
			return nil, factory.Errors.InvalidPassword()
		}
		pw, error := cipher.Hash(cmd.NewPassword)
		if error != nil {
			return nil, factory.Errors.Failed("hash")
		}
		err = repo.SetPassword(ctx, cmd.UserUUID, pw)
		if err != nil {
			return nil, err
		}
		return &ChangePasswordRes{}, nil
	}
}
