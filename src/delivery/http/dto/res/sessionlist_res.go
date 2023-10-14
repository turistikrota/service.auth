package res

import (
	"github.com/turistikrota/service.auth/src/app/query"
	"github.com/turistikrota/service.shared/auth/session"
)

type SessionListResponse struct {
	session.Entity
	IsCurrent bool `json:"is_current"`
}

func (r *response) SessionList(res *query.SessionListResult, id string) []SessionListResponse {
	list := make([]SessionListResponse, 0)
	for _, s := range res.Sessions {
		list = append(list, SessionListResponse{
			Entity:    s,
			IsCurrent: s.DeviceUUID == id,
		})
	}
	return list
}
