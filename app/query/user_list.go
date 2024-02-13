package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.auth/domains/user"
	"github.com/turistikrota/service.auth/pkg/utils"
)

type UserListQuery struct {
	*utils.Pagination
	*user.FilterEntity
}

type UserListRes struct {
	List *list.Result[user.ListDto]
}

type UserListHandler cqrs.HandlerFunc[UserListQuery, *UserListRes]

func NewUserListHandler(repo user.Repo) UserListHandler {
	return func(ctx context.Context, query UserListQuery) (*UserListRes, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		res, err := repo.List(ctx, *query.FilterEntity, list.Config{
			Offset: offset,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, err
		}
		li := make([]user.ListDto, 0, len(res.List))
		for _, u := range res.List {
			li = append(li, u.ToListDto())
		}
		return &UserListRes{
			List: &list.Result[user.ListDto]{
				List:          li,
				Total:         res.Total,
				FilteredTotal: res.FilteredTotal,
				Page:          res.Page,
				IsNext:        res.IsNext,
				IsPrev:        res.IsPrev,
			},
		}, nil
	}
}
