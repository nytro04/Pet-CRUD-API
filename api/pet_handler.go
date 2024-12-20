package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/nytro04/pet-crud/types"
)

type PetStore interface {
	CreatePet(context.Context, *types.Pet) (*types.Pet, error)
	UpdatePet(context.Context, string, *types.CreatePetParams) error
	GetPetById(context.Context, string) (*types.Pet, error)
	GetPets(context.Context) ([]*types.Pet, error)
	DeletePet(context.Context, string) (*types.Pet, error)
	Close() error
}

type PetHandler struct {
	store PetStore
}

func NewPetHandler(store PetStore) *PetHandler {
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

	pet, err := h.store.CreatePet(r.Context(), pet)
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
	pet, err := h.store.GetPetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return JSON response
	renderJSON(w, http.StatusOK, pet)
}

// GetPetsHandler gets all pets
func (h *PetHandler) GetPetsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	pets, err := h.store.GetPets(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return JSON response
	renderJSON(w, http.StatusOK, pets)
}

// UpdatePetHandler updates a pet
func (h *PetHandler) UpdatePetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.Header().Set("Allow", http.MethodPatch)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var (
		updatePayload types.CreatePetParams
		petId         = r.PathValue("id")
	)

	if err := json.NewDecoder(r.Body).Decode(&updatePayload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := h.store.GetPetById(r.Context(), petId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// @TODO: fix bug: patch does not keep the unchanged fields, it sets them to zero values

	err = h.store.UpdatePet(r.Context(), petId, &updatePayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return JSON response
	renderJSON(w, http.StatusOK, updatePayload)
}

// DeleteHandler deletes a pet
func (h *PetHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	pet, err := h.store.GetPetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.store.DeletePet(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return JSON response
	renderJSON(w, http.StatusOK, pet)
}
