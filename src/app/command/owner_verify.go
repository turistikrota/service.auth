package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.auth/src/domain/owner"
	"github.com/turistikrota/service.shared/decorator"
)

type OwnerVerifyCommand struct {
	NickName string
}

type OwnerVerifyResult struct{}

type OwnerVerifyHandler decorator.CommandHandler[OwnerVerifyCommand, *OwnerVerifyResult]

type ownerVerifyHandler struct {
	repo owner.Repository
}

type OwnerVerifyHandlerConfig struct {
	Repo     owner.Repository
	CqrsBase decorator.Base
}

func NewOwnerVerifyHandler(config OwnerVerifyHandlerConfig) OwnerVerifyHandler {
	return decorator.ApplyCommandDecorators[OwnerVerifyCommand, *OwnerVerifyResult](
		ownerVerifyHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h ownerVerifyHandler) Handle(ctx context.Context, command OwnerVerifyCommand) (*OwnerVerifyResult, *i18np.Error) {
	err := h.repo.Verify(ctx, command.NickName)
	if err != nil {
		return nil, err
	}
	return &OwnerVerifyResult{}, nil
}
