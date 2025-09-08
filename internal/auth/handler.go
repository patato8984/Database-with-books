package auth

import (
	"encoding/json"
	"log"
	"myapp/internal/models"
	"net/http"
)

type AuthHandler struct {
	AuthService *AuthService
}

func NewAuthHandler(AuthService *AuthService) *AuthHandler {
	return &AuthHandler{AuthService: AuthService}
}

func (h *AuthHandler) NewRegister(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, `{"status":"error json"}`, http.StatusBadRequest)
		return
	}
	if er := h.AuthService.Register(newUser.Name, newUser.Password); er != nil {
		http.Error(w, `{"status":"error"}`, http.StatusBadRequest)
		return
	}
}
func (h *AuthHandler) Authentication(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"status":"error json"}`, http.StatusBadRequest)
		return
	}
	jwtTocen, err := h.AuthService.GetToken(user.Name, user.Password)
	if err != nil {
		http.Error(w, `{"status":"db"}`, http.StatusBadRequest)
		log.Print(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jwtTocen)
}
