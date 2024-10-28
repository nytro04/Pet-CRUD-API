package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nytro04/pet-crud/types"
)

type AuthHandler struct {
	userStore UserStore
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

type genericResp struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

// constructor/factory function
func NewAuthHandler(userStore UserStore) *AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}

func invalidCredentials(w http.ResponseWriter) {
	renderJSON(w, http.StatusBadRequest, genericResp{
		Type: "error",
		Msg:  "invalid credentials",
	})
}

// User sign in
func (h *AuthHandler) HandleAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var params AuthParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userStore.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !types.IsValidPassword(user.EncryptedPassword, params.Password) {
		invalidCredentials(w)
		return
	}

	resp := AuthResponse{
		User:  user,
		Token: CreateTokenFromUser(user),
	}

	renderJSON(w, http.StatusOK, resp)
}

func CreateTokenFromUser(user *types.User) string {
	now := time.Now()
	expires := now.Add(time.Hour * 4).Unix()
	claims := jwt.MapClaims{
		"id":      user.ID,
		"email":   user.Email,
		"expires": expires,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("failed to sign token with secret")
	}

	return tokenStr
}
