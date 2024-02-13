package command

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.auth/config"
	"github.com/turistikrota/service.auth/domains/user"
	"github.com/turistikrota/service.auth/pkg/claims"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/cipher"
)

type LoginCmd struct {
	Email      string          `json:"email" validate:"required,email"`
	Password   string          `json:"password" validate:"required,password"`
	DeviceUUID string          `json:"-"`
	Device     *session.Device `json:"-"`
}

type LoginRes struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"-"`
}

type LoginHandler cqrs.HandlerFunc[LoginCmd, *LoginRes]

func NewLoginHandler(userRepo user.Repo, userFactory user.Factory, sessionSrv session.Service, rpc config.Rpc) LoginHandler {
	return func(ctx context.Context, cmd LoginCmd) (*LoginRes, *i18np.Error) {
		user, err := userRepo.GetByEmail(ctx, cmd.Email)
		if err != nil {
			return nil, err
		}
		if !user.IsVerified {
			return nil, userFactory.Errors.NotVerified(map[string]interface{}{
				"verifyRequired": true,
			})
		}
		if err := cipher.Compare(user.Password, cmd.Password); err != nil {
			return nil, userFactory.Errors.InvalidPassword()
		}
		accounts, businesses, error := claims.Fetch(ctx, rpc, user.UUID)
		if error != nil {
			return nil, userFactory.Errors.Failed(error.Error())
		}
		ses := &session.SessionUser{
			UUID:       user.UUID,
			Email:      user.Email,
			Roles:      user.Roles,
			Accounts:   accounts,
			Businesses: businesses,
		}
		tokens, _err := sessionSrv.New(session.NewCommand{
			UserUUID:   user.UUID,
			DeviceUUID: cmd.DeviceUUID,
			Device:     cmd.Device,
			User:       ses,
		})
		if _err != nil {
			return nil, userFactory.Errors.Failed("token")
		}
		if user.IsDeleted && time.Since(*user.DeletedAt) > 30*time.Hour*24 {
			_err := userRepo.Recover(ctx, user.UUID)
			if _err != nil {
				return nil, _err
			}
		}
		return &LoginRes{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		}, nil
	}
}
