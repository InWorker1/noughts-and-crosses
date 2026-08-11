package auth

type AuthService interface {
	Register(request SignUpRequest) error
	Login(request SignUpRequest) (string, error)
}

type AuthRepository interface {
	Create(request SignUpRequest) error
	GetByUsername(username string) (string, error)
}
