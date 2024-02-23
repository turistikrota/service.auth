package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.shared/auth/session"
)

type SessionListQuery struct {
	UserUUID   string `json:"-"`
	DeviceUUID string `json:"-"`
}

type SessionListItem struct {
	session.Entity
	IsCurrent bool `json:"is_current"`
}

type SessionListRes struct {
	Sessions []SessionListItem
}

type SessionListHandler cqrs.HandlerFunc[SessionListQuery, *SessionListRes]

func NewSessionListHandler(sessionSrv session.Service) SessionListHandler {
	return func(ctx context.Context, query SessionListQuery) (*SessionListRes, *i18np.Error) {
		sessions, err := sessionSrv.GetAll(ctx, query.UserUUID)
		if err != nil {
			return nil, err
		}
		res := make([]SessionListItem, 0, len(sessions))
		for _, s := range sessions {
			res = append(res, SessionListItem{
				Entity:    s,
				IsCurrent: s.DeviceUUID == query.DeviceUUID,
			})
		}
		return &SessionListRes{
			Sessions: res,
		}, nil
	}
}
