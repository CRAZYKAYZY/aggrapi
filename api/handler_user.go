package api

import (
	"encoding/json"
	"net/http"
	"time"

	sqlc "github.com/CRAZYKAYZY/aggrapi/db/sqlc"
	"github.com/google/uuid"
)

func (server *Server) HandlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type Users struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	userparams := Users{}
	err := decoder.Decode(&userparams)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "could not decode")
		return
	}

	uuidValue := uuid.New()
	user, err := server.store.CreateUser(r.Context(), sqlc.CreateUserParams{
		ID:        int32(uuidValue.ID()),
		Name:      userparams.Name,
		Email:     userparams.Email,
		Password:  userparams.Password,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	RespondWithJSON(w, http.StatusOK, user)
}
