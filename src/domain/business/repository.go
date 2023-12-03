package business

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/jwt"
)

type UserDetail struct {
	Name string
	Code string
	UUID string
}

type Repository interface {
	Create(ctx context.Context, entity *Entity) *i18np.Error
	AddUser(ctx context.Context, nickName string, user *User) *i18np.Error
	RemoveUser(ctx context.Context, nickName string, user UserDetail) *i18np.Error
	RemoveUserPermission(ctx context.Context, nickName string, user UserDetail, permission string) *i18np.Error
	AddUserPermission(ctx context.Context, nickName string, user UserDetail, permission string) *i18np.Error
	Enable(ctx context.Context, nickName string) *i18np.Error
	Disable(ctx context.Context, nickName string) *i18np.Error
	Delete(ctx context.Context, nickName string) *i18np.Error
	Recover(ctx context.Context, nickName string) *i18np.Error
	Verify(ctx context.Context, nickName string) *i18np.Error
	GetAllAsClaim(ctx context.Context, userUUID string) ([]jwt.UserClaimBusiness, *i18np.Error)
}
