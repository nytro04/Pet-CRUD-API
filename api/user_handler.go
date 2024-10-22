package api

import (
	"encoding/json"
	"net/http"

	"github.com/nytro04/pet-crud/db"
	"github.com/nytro04/pet-crud/types"
)

type UserHandler struct {
	store *db.Store
}

// constructor/factory function
func NewUserHandler(store *db.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

// add new comment
// User sign up
func (h *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var params types.CreateUserParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &types.User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: params.Password,
	}

	insertedUser, err := h.store.User.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return JSON response
	renderJSON(w, http.StatusCreated, insertedUser)
}

// Get user by ID
func (h *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	user, err := h.store.User.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	renderJSON(w, http.StatusOK, user)
}

func (h *UserHandler) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.User.GetUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(w, http.StatusOK, users)
}

// TODO: Update single property without losing the other properties typical PATCH, not PUT
func (h *UserHandler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var (
		params *types.UpdateUserParams
		userID = r.PathValue("id")
	)

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := h.store.User.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.store.User.UpdateUser(r.Context(), userID, params); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responsePayload := types.User{
		ID:        userID,
		FirstName: params.FirstName,
		LastName:  params.LastName,
	}

	renderJSON(w, http.StatusOK, responsePayload)
}

func (h *UserHandler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")

	_, err := h.store.User.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = h.store.User.DeleteUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSON(w, http.StatusOK, nil)
}
