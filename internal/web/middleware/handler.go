package middleware

import (
	"game/internal/domain/auth"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type UserAuthenticator struct {
	authService auth.AuthService
}

func NewUserAuthenticator(s auth.AuthService) *UserAuthenticator {
	return &UserAuthenticator{authService: s}
}

func (m *UserAuthenticator) Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cred := r.Header.Get("Authorization")
		cred = strings.TrimPrefix(cred, "Basic ")
		cred = strings.TrimSpace(cred)
		if cred == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		id, err := m.authService.Login(cred)
		if id == uuid.Nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, "sorry", http.StatusInternalServerError)
			return
		}

		next(w, r)
	}

}
