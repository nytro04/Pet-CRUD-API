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

func (h *PetHandler) CreatePetHandler(c *fiber.Ctx) error {
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

func (h *PetHandler) GetPetHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	pet, err := h.store.Pet.GetPet(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(pet)
}

func (h *PetHandler) GetPetsHandler(c *fiber.Ctx) error {
	pets, err := h.store.Pet.GetPets(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(pets)
}

func (h *PetHandler) UpdatePetHandler(c *fiber.Ctx) error {
	var (
		updatePayload *types.CreatePetParams
		petId         = c.Params("id")
	)

	if err := c.BodyParser(&updatePayload); err != nil {
		return err
	}

	err := h.store.Pet.UpdatePet(c.Context(), petId, updatePayload)
	if err != nil {
		return err
	}

	responsePayload := types.Pet{
		ID:    petId,
		Name:  updatePayload.Name,
		Owner: updatePayload.Owner,
		Type:  updatePayload.Type,
		Age:   updatePayload.Age,
	}

	return c.JSON(responsePayload)
}
