package api

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nytro04/pet-crud/db"
	"github.com/nytro04/pet-crud/types"
)

type UserHandler struct {
	userStore *db.Store
}

// constructor/factory function
func NewUserHandler(userStore *db.Store) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

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

	insertedUser, err := h.userStore.User.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return JSON response
	renderJSON(w, http.StatusCreated, insertedUser)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)

	user, err := h.userStore.User.GetUserByID(c.Context(), id)
	if err != nil {

		return err

	}

	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.User.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	var (
		// values bson.E
		params *types.UpdateUserParams
		userID = c.Params("id")
	)

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	_, err := h.userStore.User.GetUserByID(c.Context(), userID)
	if err != nil {
		return err
	}

	if err := h.userStore.User.UpdateUser(c.Context(), userID, params); err != nil {
		return err
	}

	responsePayload := types.User{
		ID:        userID,
		FirstName: params.FirstName,
		LastName:  params.LastName,
	}

	return c.JSON(responsePayload)
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	user, err := h.userStore.User.GetUserByID(c.Context(), userID)
	if err != nil {
		return err
	}

	_, err = h.userStore.User.DeleteUser(c.Context(), userID)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
