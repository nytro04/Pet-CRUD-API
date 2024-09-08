package api

import (
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
func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if errors := params.Validate(); len(errors) > 0 {
		return c.JSON(errors)
	}
	user := &types.User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: params.Password,
	}

	insertedUser, err := h.userStore.User.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(insertedUser)
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
