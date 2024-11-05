package db

import (
	"context"
	"testing"
	"time"

	"github.com/nytro04/pet-crud/mocks"
	"github.com/nytro04/pet-crud/types"
	"github.com/stretchr/testify/assert"
)

// unit test for GetPetById using mocks from ../mocks/PetStore.go
func TestGetPetById(t *testing.T) {
	// create a mock object
	mock := mocks.MockPetStore(t)
	// create a new PetStorage object with the mock object
	// create a new context
	ctx := context.Background()
	// create a new pet object
	pet := &types.Pet{
		ID:   "1",
		Name: "dog",
	}
	// mock the GetPetById function
	mock.On("GetPetById", ctx, "1").Return(pet, nil)
	// call the GetPetById function
	result, err := mock.GetPetById(ctx, "1")
	// assert the result
	assert.Equal(t, pet, result)
	// assert the error
	assert.Nil(t, err)
	// assert that the mock was called
	mock.AssertExpectations(t)

}

// unit test for CreatePet using mocks from ../mocks/PetStore.go
func TestCreatePet(t *testing.T) {
	// create a mock object
	mock := mocks.MockPetStore(t)
	// create a new PetStorage object with the mock object
	// create a new context
	ctx := context.Background()
	// create a new pet object
	pet := &types.Pet{
		ID:        "1",
		Name:      "dog",
		Age:       2,
		Owner:     "owner",
		Type:      "dog",
		CreatedAt: time.Now().UTC(),
	}
	// mock the CreatePet function
	mock.On("CreatePet", ctx, pet).Return(pet, nil)
	// call the CreatePet function
	result, err := mock.CreatePet(ctx, pet)
	// assert the result
	assert.Equal(t, pet, result)
	// assert the error
	assert.Nil(t, err)
	// assert that the mock was called
	mock.AssertExpectations(t)
}

// unit test for GetPets using mocks from ../mocks/PetStore.go
func TestGetPets(t *testing.T) {
	// create a mock object
	mock := mocks.MockPetStore(t)
	// create a new PetStorage object with the mock object
	// create a new context
	ctx := context.Background()
	// create a new pet object
	pet := &types.Pet{
		ID:        "1",
		Name:      "dog",
		Age:       2,
		Owner:     "owner",
		Type:      "dog",
		CreatedAt: time.Now().UTC(),
	}
	// mock the GetPets function
	mock.On("GetPets", ctx).Return([]*types.Pet{pet}, nil)
	// call the GetPets function
	result, err := mock.GetPets(ctx)
	// assert the result
	assert.Equal(t, []*types.Pet{pet}, result)
	// assert the error
	assert.Nil(t, err)
	// assert that the mock was called
	mock.AssertExpectations(t)
}

// unit test for UpdatePet using mocks from ../mocks/PetStore.go
func TestUpdatePet(t *testing.T) {
	// create a mock object
	mock := mocks.MockPetStore(t)
	// create a new PetStorage object with the mock object
	// create a new context
	ctx := context.Background()

	// mock the UpdatePet function
	mock.On("UpdatePet", ctx, "1", &types.CreatePetParams{Name: "kiwi"}).Return(nil)
	// call the UpdatePet function
	err := mock.UpdatePet(ctx, "1", &types.CreatePetParams{Name: "kiwi"})
	// assert the error
	assert.Nil(t, err)
	// assert that the mock was called
	mock.AssertExpectations(t)
}

// unit test for DeletePet using mocks from ../mocks/PetStore.go
func TestDeletePet(t *testing.T) {
	// create a mock object
	mock := mocks.MockPetStore(t)
	// create a new PetStorage object with the mock object
	// create a new context
	ctx := context.Background()
	// create a new pet object
	pet := &types.Pet{
		ID:        "1",
		Name:      "dog",
		Age:       2,
		Owner:     "owner",
		Type:      "dog",
		CreatedAt: time.Now().UTC(),
	}
	// mock the DeletePet function
	mock.On("DeletePet", ctx, "1").Return(pet, nil)
	// call the DeletePet function
	result, err := mock.DeletePet(ctx, "1")
	// assert the result
	assert.Equal(t, pet, result)
	// assert the error
	assert.Nil(t, err)
	// assert that the mock was called
	mock.AssertExpectations(t)
}
