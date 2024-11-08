package api

import (
	"encoding/json"
	"net/http"

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

func NewUserHandler(service services.UserService, jwtSecret string) *UserHandler {
	return &UserHandler{
		UserService: service,
		JWTSecret:   jwtSecret,
	}
}

func (u *UserHandler) RegisterRoutes(r chi.Router) {
	r.Post("/new-user", u.HandleNewUser)
	r.Post("/login", u.Login)
	//r.Put("/update-user", u.UpdateUserHandler)
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

// // Update User handler
// func (server *Server) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
// 	// Extract the authorization token from the request header
// 	tokenString := r.Header.Get("Authorization")
// 	if tokenString == "" {
// 		util.RespondWithError(w, http.StatusUnauthorized, "missing authorization token")
// 		return
// 	}

// 	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

// 	// Parse and validate the authorization token
// 	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		// Return the secret key used to sign the token
// 		jwtSecret := os.Getenv("JWT_SECRET")
// 		return []byte(jwtSecret), nil
// 	})
// 	if err != nil {
// 		util.RespondWithError(w, http.StatusUnauthorized, "invalid authorization token")
// 		return
// 	}

// 	// Verify if the token is valid and contains the required claims
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		util.RespondWithError(w, http.StatusUnauthorized, "failed to verify")
// 		return
// 	}

// 	// Extract the user ID from the claims
// 	id, ok := claims["sub"].(string)
// 	if !ok {
// 		util.RespondWithError(w, http.StatusUnauthorized, "invalid user id")
// 		return
// 	}

// 	userID, err := uuid.Parse(id)
// 	if err != nil {
// 		util.RespondWithError(w, http.StatusBadRequest, "invalid user ID")
// 		return
// 	}

// 	// Get the user from the database using the id
// 	user, err := server.store.GetUser(r.Context(), sqlc.GetUserParams{ID: userID})
// 	if err != nil {
// 		// Handle the error, e.g., log it, return an error response, etc.
// 		log.Printf("Error getting user: %v", err)
// 		util.RespondWithError(w, http.StatusInternalServerError, "Failed to get user")
// 		return
// 	}
// 	// Implement logic to update the user fields like name and email

// 	type UpdateUserReq struct {
// 		Name     string `json:"name"`
// 		Email    string `json:"email"`
// 		Password string `json:"password,omitempty"`
// 	}

// 	var updateReq UpdateUserReq
// 	err = json.NewDecoder(r.Body).Decode(&updateReq)
// 	if err != nil {
// 		util.RespondWithError(w, http.StatusBadRequest, "failed to decode request body")
// 		return
// 	}

// 	// Update the user fields only if they are not empty
// 	if updateReq.Name != "" {
// 		user.Name = updateReq.Name
// 	}
// 	if updateReq.Email != "" {
// 		user.Email = updateReq.Email
// 	}

// 	// Only update the password if it's not empty
// 	if updateReq.Password != "" {
// 		hashedPassword, err := util.HashedPass(updateReq.Password)
// 		if err != nil {
// 			util.RespondWithError(w, http.StatusInternalServerError, "couldn't hash password")
// 			return
// 		}
// 		user.Password = hashedPassword
// 	}

// 	// Create the UpdateUserParams with the updated values
// 	updateParams := sqlc.UpdateUserParams{
// 		ID:       user.ID,
// 		Name:     user.Name,
// 		Email:    user.Email,
// 		Password: user.Password,
// 	}

// 	// if err != nil {
// 	// 	// Handle the error, e.g., log it, return an error response, etc.
// 	// 	log.Printf("Error creating updateParams: %v", err)
// 	// 	RespondWithError(w, http.StatusInternalServerError, "Failed to create updateParams")
// 	// 	return
// 	// }

// 	updatedUser, err := server.store.UpdateUser(r.Context(), updateParams)

// 	if err != nil {
// 		// Handle the error, e.g., log it, return an error response, etc.
// 		log.Printf("Error creating updateParams: %v", err)
// 		util.RespondWithError(w, http.StatusInternalServerError, "Failed to update user")
// 		return
// 	}

// 	// Return a success response
// 	util.RespondWithJSON(w, http.StatusOK, updatedUser)
// }
