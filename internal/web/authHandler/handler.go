package authHandler

import (
	"encoding/json"
	"game/internal/domain/auth"
	"net/http"
	"strings"
)

type AuthHandler struct {
	service auth.AuthService
}

func NewAuthHandler(service auth.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req JsonRequestReg
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	domainReq := reqIntoDomain(req)
	err = h.service.Register(domainReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	cred := r.Header.Get("Authorization")
	if cred == "" {
		http.Error(w, "invalid authorization header", http.StatusUnauthorized)
		return
	}
	cred = strings.TrimPrefix(cred, "Basic ")
	cred = strings.TrimSpace(cred)
	uuid, err := h.service.Login(cred)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("id", uuid.String())

	w.WriteHeader(http.StatusOK)
}
