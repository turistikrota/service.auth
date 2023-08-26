package user

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/types/list"
)

type Repository interface {
	GetByUUID(ctx context.Context, uuid string) (*User, *i18np.Error)
	GetByEmail(ctx context.Context, email string) (*User, *i18np.Error)
	GetByToken(ctx context.Context, token string) (*User, *i18np.Error)
	CheckEmail(ctx context.Context, email string) (bool, *i18np.Error)
	Create(ctx context.Context, email string, password []byte, token string) (*User, *i18np.Error)
	Update(ctx context.Context, user *User) (*User, *i18np.Error)
	SetToken(ctx context.Context, email string, token string) *i18np.Error
	Verify(ctx context.Context, token string) *i18np.Error
	UpdateByUUID(ctx context.Context, user *User) (*User, *i18np.Error)
	List(ctx context.Context, config list.Config) (*list.Result[*ListEntity], *i18np.Error)
	AddRoles(ctx context.Context, uuid string, roles []string) *i18np.Error
	RemoveRoles(ctx context.Context, uuid string, roles []string) *i18np.Error
}
