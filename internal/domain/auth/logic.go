package auth

import (
	"game/internal/domain/domainErrors"
	"game/internal/domain/user"

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
	id := a.uService.Register(request.Login, request.Pass)
	if id == uuid.Nil {
		return domainErrors.ErrInvalidLoginOrPass
	}
	return nil
} // гуд

func (a *authService) Login(request SignUpRequest) (string, error) {
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
