package query

import (
	"context"
	"strconv"
	"time"

	"github.com/mixarchitecture/cache"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/mixarchitecture/microp/types/list"
	"github.com/turistikrota/service.auth/src/domain/user"
)

type UserListQuery struct {
	Offset int64
	Limit  int64
}

type UserListResult struct {
	List *list.Result[*user.ListEntity]
}

type UserListHandler decorator.QueryHandler[UserListQuery, *UserListResult]

type userListHandler struct {
	repo  user.Repository
	cache cache.Client[*list.Result[*user.ListEntity]]
}

type UserListHandlerConfig struct {
	Repo     user.Repository
	CqrsBase decorator.Base
	CacheSrv cache.Service
}

func NewUserListHandler(config UserListHandlerConfig) UserListHandler {
	return decorator.ApplyQueryDecorators[UserListQuery, *UserListResult](
		userListHandler{
			repo:  config.Repo,
			cache: cache.New[*list.Result[*user.ListEntity]](config.CacheSrv),
		},
		config.CqrsBase,
	)
}

func (h userListHandler) Handle(ctx context.Context, query UserListQuery) (*UserListResult, *i18np.Error) {
	cacheHandler := func() (*list.Result[*user.ListEntity], *i18np.Error) {
		return h.repo.List(ctx, list.Config{
			Offset: query.Offset,
			Limit:  query.Limit,
		})
	}
	res, err := h.cache.Creator(h.createCacheEntity).Handler(cacheHandler).Timeout(1*time.Minute).Get(ctx, h.generateCacheKey(query))
	if err != nil {
		return nil, err
	}
	return &UserListResult{
		List: res,
	}, nil
}

func (h userListHandler) createCacheEntity() *list.Result[*user.ListEntity] {
	return &list.Result[*user.ListEntity]{}
}

func (h userListHandler) generateCacheKey(query UserListQuery) string {
	return "list_users_" + strconv.FormatInt(query.Offset, 10) + "_" + strconv.FormatInt(query.Limit, 10)
}
