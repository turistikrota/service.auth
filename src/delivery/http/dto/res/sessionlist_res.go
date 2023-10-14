package res

import (
	"github.com/turistikrota/service.auth/src/app/query"
	"github.com/turistikrota/service.shared/auth/session"
)

func (r *response) SessionList(res *query.SessionListResult) []session.Entity {
	return res.Sessions
}
