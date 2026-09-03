package middleware

import "game/internal/domain/auth"

type middlewareStruct struct {
	authService auth.AuthService
}

func NewMiddleWare(s auth.AuthService) *middlewareStruct {
	return &middlewareStruct{authService: s}
}

func (*middlewareStruct) Authorization() {

}
