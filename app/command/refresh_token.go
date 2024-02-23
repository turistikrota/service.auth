package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.auth/config"
	"github.com/turistikrota/service.auth/domains/user"
	"github.com/turistikrota/service.auth/pkg/claims"
	"github.com/turistikrota/service.shared/auth/session"
)

type RefreshTokenCmd struct {
	RefreshToken string
	AccessToken  string
	DeviceUUID   string
	UserUUID     string
	IpAddress    string
}

type RefreshTokenRes struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"-"`
}

type RefreshTokenHandler cqrs.HandlerFunc[RefreshTokenCmd, *RefreshTokenRes]

func NewRefreshTokenHandler(sessionSrv session.Service, repo user.Repo, factory user.Factory, rpc config.Rpc) RefreshTokenHandler {
	return func(ctx context.Context, cmd RefreshTokenCmd) (*RefreshTokenRes, *i18np.Error) {
		available := sessionSrv.IsRefreshAvailable(ctx, session.IsRefreshAvailableCommand{
			UserUUID:     cmd.UserUUID,
			DeviceUUID:   cmd.DeviceUUID,
			AccessToken:  cmd.AccessToken,
			RefreshToken: cmd.RefreshToken,
		})
		if !available {
			return nil, factory.Errors.RefreshTokenNotAvailable()
		}
		user, err := repo.GetByUUID(ctx, cmd.UserUUID)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		accounts, business, error := claims.Fetch(ctx, rpc, user.UUID)
		if error != nil {
			return nil, factory.Errors.AnErrorOccurred()
		}
		ses := &session.SessionUser{
			UUID:             user.UUID,
			Email:            user.Email,
			Phone:            user.Phone,
			Roles:            user.Roles,
			Accounts:         accounts,
			Businesses:       business,
			TwoFactorEnabled: user.TwoFactorEnabled,
		}
		tokens, err := sessionSrv.Refresh(ctx, session.RefreshCommand{
			UserUUID:     cmd.UserUUID,
			DeviceUUID:   cmd.DeviceUUID,
			RefreshToken: cmd.RefreshToken,
			AccessToken:  cmd.AccessToken,
			User:         ses,
			IpAddress:    cmd.IpAddress,
		})
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &RefreshTokenRes{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		}, nil
	}
}
