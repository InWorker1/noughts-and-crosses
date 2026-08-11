package auth

type SignUpRequest struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}
