package res

type VerifyRequiredResponse struct {
	Verify bool `json:"verify"`
}

func (r *response) VerifyRequired() *VerifyRequiredResponse {
	return &VerifyRequiredResponse{
		Verify: true,
	}
}
