package res

type AuthResponse struct {
	AccessToken string `json:"token"`
}

func (r *response) LoggedIn(token string) *AuthResponse {
	return &AuthResponse{
		AccessToken: token,
	}
}
