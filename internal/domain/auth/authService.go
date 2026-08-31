package auth

import "github.com/google/uuid"

type AuthService interface {
	Register(request SignUpRequest) error
	Login(creds string) (uuid.UUID, error)
}
