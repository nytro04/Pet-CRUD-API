package api

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nytro04/pet-crud/db"
	"github.com/nytro04/pet-crud/types"
)

type PetHandler struct {
	store *db.Store
}

func NewPetHandler(store *db.Store) *PetHandler {
	return &PetHandler{
		store: store,
	}
}

func (h *PetHandler) HandleCreatePet(c *fiber.Ctx) error {
	var params types.CreatePetParams

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	pet := &types.Pet{
		Name:      params.Name,
		Age:       params.Age,
		Owner:     params.Owner,
		Type:      params.Type,
		CreatedAt: time.Now().UTC(),
	}

	pet, err := h.store.Pet.CreatePet(c.Context(), pet)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(pet)
}
