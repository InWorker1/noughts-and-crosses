package middleware

import (
	"game/internal/domain/auth"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type middlewareStruct struct {
	authService auth.AuthService
}

func NewMiddleWare(s auth.AuthService) *middlewareStruct {
	return &middlewareStruct{authService: s}
}

func (m *middlewareStruct) Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cred := r.Header.Get("Authorization")
		cred = strings.TrimPrefix(cred, "Bearer ")
		cred = strings.TrimSpace(cred)
		if cred != "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		id, err := m.authService.Login(cred)
		if id != uuid.Nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		} else if err != nil {
			http.Error(w, "sorry", http.StatusInternalServerError)
		}

		next(w, r)
	}

}
