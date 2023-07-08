package query

import (
	"context"
	"time"

	"github.com/mixarchitecture/cache"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.auth/src/domain/user"
)

type CheckEmailQuery struct {
	Email string
}

type CheckEmailResult struct {
	Exists bool
}

type CheckEmailHandler decorator.QueryHandler[CheckEmailQuery, *CheckEmailResult]

type checkEmailHandler struct {
	repo  user.Repository
	cache cache.Client[bool]
}

type CheckEmailHandlerConfig struct {
	Repo     user.Repository
	CqrsBase decorator.Base
	CacheSrv cache.Service
}

func NewCheckEmailHandler(config CheckEmailHandlerConfig) CheckEmailHandler {
	return decorator.ApplyQueryDecorators[CheckEmailQuery, *CheckEmailResult](
		checkEmailHandler{
			repo:  config.Repo,
			cache: cache.New[bool](config.CacheSrv),
		},
		config.CqrsBase,
	)
}

func (h checkEmailHandler) Handle(ctx context.Context, query CheckEmailQuery) (*CheckEmailResult, *i18np.Error) {
	cacheHandler := func() (bool, *i18np.Error) {
		return h.repo.CheckEmail(ctx, query.Email)
	}
	res, err := h.cache.Creator(h.createCacheEntity).Handler(cacheHandler).Timeout(1*time.Minute).Get(ctx, h.generateCacheKey(query))
	if err != nil {
		return nil, err
	}
	return &CheckEmailResult{
		Exists: res,
	}, nil
}

func (h checkEmailHandler) createCacheEntity() bool {
	return false
}

func (h checkEmailHandler) generateCacheKey(query CheckEmailQuery) string {
	return "check_email_" + query.Email
}
