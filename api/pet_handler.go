package api

import (
	"encoding/json"
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

// CreatePetHandler creates a new pet
func (h *PetHandler) CreatePetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var params types.CreatePetParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pet := &types.Pet{
		Name:      params.Name,
		Age:       params.Age,
		Owner:     params.Owner,
		Type:      params.Type,
		CreatedAt: time.Now().UTC(),
	}

	pet, err := h.store.Pet.CreatePet(r.Context(), pet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return JSON response
	renderJSON(w, http.StatusCreated, pet)
}

// GetPetByIdHandler gets a pet by its ID
func (h *PetHandler) GetPetByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	pet, err := h.store.Pet.GetPetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return JSON response
	renderJSON(w, http.StatusOK, pet)
}

// func (h *PetHandler) GetPetByIdHandler(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	pet, err := h.store.Pet.GetPetById(c.Context(), id)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(pet)
// }

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

	// @TODO: add ternary like to use pet or payload
	// pet, err := h.store.Pet.GetPetById(c.Context(), petId)
	_, err := h.store.Pet.GetPetById(c.Context(), petId)
	if err != nil {
		return err
	}

	err = h.store.Pet.UpdatePet(c.Context(), petId, updatePayload)
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

func (h *PetHandler) DeleteHandler(c *fiber.Ctx) error {
	petId := c.Params("id")
	pet, err := h.store.Pet.GetPetById(c.Context(), petId)
	if err != nil {
		return err
	}

	_, err = h.store.Pet.DeletePet(c.Context(), petId)
	if err != nil {
		return err
	}

	return c.JSON(pet)
}
