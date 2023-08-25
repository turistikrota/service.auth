package res

import (
	"github.com/mixarchitecture/microp/types/list"
	"github.com/turistikrota/service.auth/src/app/query"
	"github.com/turistikrota/service.auth/src/domain/user"
)

func (r *response) UserList(res *query.UserListResult) *list.Result[*user.ListEntity] {
	return res.List
}
