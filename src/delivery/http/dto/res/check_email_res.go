package res

import "api.turistikrota.com/auth/src/app/query"

type CheckEmailResponse struct {
	Exists bool `json:"exists"`
}

func (r *response) CheckEmail(result *query.CheckEmailResult) *CheckEmailResponse {
	return &CheckEmailResponse{
		Exists: result.Exists,
	}
}
