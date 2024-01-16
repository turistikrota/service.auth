package account

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/jwt"
)

type UserUnique struct {
	UserUUID string
	Name     string
}

type Repository interface {
	Create(ctx context.Context, account *Entity) *i18np.Error
	Update(ctx context.Context, u UserUnique, account *Entity) *i18np.Error
	Disable(ctx context.Context, u UserUnique) *i18np.Error
	Enable(ctx context.Context, u UserUnique) *i18np.Error
	Restore(ctx context.Context, u UserUnique) *i18np.Error
	Delete(ctx context.Context, u UserUnique) *i18np.Error
	ListAsClaims(ctx context.Context, userUUID string) ([]jwt.UserClaimAccount, *i18np.Error)
}
