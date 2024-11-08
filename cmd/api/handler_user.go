package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ChileKasoka/mis/internal/models"
	"github.com/ChileKasoka/mis/internal/services"
	"github.com/ChileKasoka/mis/middleware"

	//"github.com/ChileKasoka/mis/util"
	"github.com/go-chi/chi"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	UserService services.UserService
	JWTSecret   string
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

func NewUserHandler(service services.UserService, jwtSecret string) *UserHandler {
	return &UserHandler{
		UserService: service,
		JWTSecret:   jwtSecret,
	}
}

func (u *UserHandler) RegisterRoutes(r chi.Router) {
	r.Post("/new-user", u.HandleNewUser)
	r.Post("/login", u.Login)
	r.Post("/refresh-token", u.RefreshTokenHandler)
	r.With(middleware.JWTAuth(u.JWTSecret)).Put("/update-user", u.UpdateUserHandler)
}

func (u *UserHandler) HandleNewUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newUser, err := u.UserService.PostNewUser(user.Name, user.Email, user.Password, string(user.UserType))
	if err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := u.UserService.LoginService(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	res := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (u *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	var req UpdateUserRequest
	// Parse JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "Unauthorized: missing or invalid user ID", http.StatusUnauthorized)
		return
	}

	// Call the service to update user
	updatedUser, err := u.UserService.UpdateUserService(userID, req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the updated user details
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (u *UserHandler) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header is required", http.StatusUnauthorized)
		return
	}

	// Trim "Bearer " prefix from the token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		http.Error(w, "Bearer token required", http.StatusUnauthorized)
		return
	}

	newAccessToken, err := u.UserService.RefreshTokenService(tokenString, u.JWTSecret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Respond with the new access token
	res := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: newAccessToken,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
