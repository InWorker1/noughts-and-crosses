package auth

type AuthService interface {
	Register(request SignUpRequest) error
	Login(request SignUpRequest) (string, error)
}
