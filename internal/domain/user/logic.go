package user

import (
	"errors"
	"game/internal/domain/domainErrors"

	"github.com/google/uuid"
)

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) SaveNewPerson(login, password string) uuid.UUID {
	_, err := s.repo.GetByUsername(login)
	if !errors.Is(err, domainErrors.ErrPersonNotFound) {
		return uuid.Nil
	}
	id := uuid.New()
	user := User{Id: id,
		Login: login,
		Pass:  password,
	}
	err = s.repo.Create(user)
	if err != nil {
		return uuid.Nil
	}
	return id
}

func (s *userService) GetPerson(login, password string) (User, error) {
	user, err := s.repo.GetByUsername(login)
	if err != nil {
		return User{}, domainErrors.ErrPersonNotFound
	}
	return user, nil
}

//func (s *userService) GetRole(id uuid.UUID) (string, error) {
//	user, err := s.repo.GetByUsername(login)
//	if err != nil {
//		return "", err
//	}
//	return user, nil
//}
