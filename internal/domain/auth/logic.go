package auth

import (
	"encoding/base64"
	"game/internal/domain/domainErrors"
	"game/internal/domain/user"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	uService user.UserService
}

func NewAuthService(r user.UserService) AuthService {
	return &authService{uService: r}
}

func (a *authService) Register(request SignUpRequest) error {
	var err error
	request.Pass, err = HashPassword(request.Pass)
	if err != nil {
		return err
	}
	id := a.uService.SaveNewPerson(request.Login, request.Pass)
	if id == uuid.Nil {
		return domainErrors.ErrInvalidLoginOrPass
	}
	return nil
} // гуд

func (a *authService) Login(creds string) (uuid.UUID, error) {
	decCreds, err := base64.StdEncoding.DecodeString(creds)
	if err != nil {
		return uuid.Nil, err
	}
	parts := strings.SplitN(creds, ":", 2)
	if len(parts) != 2 {
		return uuid.Nil, domainErrors.ErrInvalidLoginOrPass
	}

	user, err := a.uService.GetPerson(parts[0], parts[1])
	if err != nil {
		return uuid.Nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(parts[1]))
	if err != nil {
		return uuid.Nil, domainErrors.ErrInvalidLoginOrPass
	}

	return user.Id, nil

	//passDB, err := a.uService.GetByUsername(request.Login)
	//if err != nil {
	//	return "", err
	//}
	//
	//err = bcrypt.CompareHashAndPassword([]byte(passDB), []byte(request.Pass))

	//if err != nil {
	//	return "", errors.New("invalid password")
	//}
	//
	//token := fmt.Sprintf("%s:%s", request.Login, request.Pass)
	//token = base64.StdEncoding.EncodeToString([]byte(token))
	//
	//return token, nil
}

func HashPassword(password string) (string, error) {
	hashB, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashB), nil
}
