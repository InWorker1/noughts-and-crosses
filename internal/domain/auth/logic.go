package auth

import (
	"errors"
	"game/internal/domain/domainErrors"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo AuthRepository
}

func NewAuthService(r AuthRepository) AuthService {
	return &authService{repo: r}
}

func (a *authService) Register(request SignUpRequest) error {
	_, err := a.repo.GetByUsername(request.Login)
	if !errors.Is(err, domainErrors.ErrPersonNotFound) {
		return err
	}

	hashB, err := bcrypt.GenerateFromPassword([]byte(request.Pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	request.Pass = string(hashB)

	err = a.repo.Create(request)
	if err != nil {
		return err
	}

	return nil
}

func (a *authService) Login(request SignUpRequest) (string, error) {
	return "", nil
}
