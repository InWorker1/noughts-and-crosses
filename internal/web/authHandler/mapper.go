package authHandler

import "game/internal/domain/auth"

func domainIntoReq(dom auth.SignUpRequest) JsonRequestReg {
	return JsonRequestReg{Login: dom.Login, Password: dom.Pass}
}

func reqIntoDomain(req JsonRequestReg) auth.SignUpRequest {
	return auth.SignUpRequest{Login: req.Login, Pass: req.Password}
}
