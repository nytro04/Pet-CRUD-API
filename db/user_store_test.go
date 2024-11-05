package db

import (
	"context"
	"testing"

	"github.com/nytro04/pet-crud/mocks"
	"github.com/nytro04/pet-crud/types"
	"github.com/stretchr/testify/assert"
)

// unit test for CreateUser using mocks from ../mocks/UserStore.go
func TestCreateUser(t *testing.T) {
	// create a mock object
	mock := mocks.NewUserStore(t)
	// create a new UserStorage object with the mock object
	// create a new context
	ctx := context.Background()
	// create a new user object
	user := &types.User{
		ID:                "1",
		FirstName:         "John",
		LastName:          "Doe",
		Email:             "john@gmail.com",
		EncryptedPassword: "password",
	}
	// mock the CreateUser function
	mock.On("CreateUser", ctx, user).Return(user, nil)
	// call the CreateUser function
	result, err := mock.CreateUser(ctx, user)
	// assert the result
	assert.Equal(t, user, result)
	// assert the error
	assert.Nil(t, err)
	// assert that the mock was called
	mock.AssertExpectations(t)

}

// unit test for GetUserByEmail using mocks from ../mocks/UserStore.go
func TestGetUserByEmail(t *testing.T) {
	// create a mock object
	mock := mocks.NewUserStore(t)
	// create a new UserStorage object with the mock object
	// create a new context
	ctx := context.Background()
	// create a new user object
	user := &types.User{
		ID:                "1",
		FirstName:         "John",
		LastName:          "Doe",
		Email:             "john@gmail.com",
		EncryptedPassword: "password",
	}
	// mock the GetUserByEmail function
	mock.On("GetUserByEmail", ctx, "john@gmail.com").Return(user, nil)
	// call the GetUserByEmail function
	result, err := mock.GetUserByEmail(ctx, "john@gmail.com")
	// assert the result
	assert.Equal(t, user, result)
	// assert the error
	assert.Nil(t, err)
	// assert that the mock was called
	mock.AssertExpectations(t)

}

// unit test for GetUserByID using mocks from ../mocks/UserStore.go
func TestGetUserByID(t *testing.T) {
	// create a mock object
	mock := mocks.NewUserStore(t)
	// create a new UserStorage object with the mock object
	// create a new context
	ctx := context.Background()
	// create a new user object
	user := &types.User{
		ID:                "1",
		FirstName:         "John",
		LastName:          "Doe",
		Email:             "john@gmail.com",
		EncryptedPassword: "password",
	}
	// mock the GetUserByID function
	mock.On("GetUserByID", ctx, "1").Return(user, nil)
	// call the GetUserByID function
	result, err := mock.GetUserByID(ctx, "1")
	// assert the result
	assert.Equal(t, user, result)
	// assert the error
	assert.Nil(t, err)
	// assert that the mock was called
	mock.AssertExpectations(t)

}

// unit test for GetUsers using mocks from ../mocks/UserStore.go
func TestGetUsers(t *testing.T) {
	// create a mock object
	mock := mocks.NewUserStore(t)
	// create a new UserStorage object with the mock object
	// create a new context
	ctx := context.Background()
	// create a new user object
	user := &types.User{
		ID:                "1",
		FirstName:         "John",
		LastName:          "Doe",
		Email:             "john@gmail.com",
		EncryptedPassword: "password",
	}

	// mock the GetUsers function
	mock.On("GetUsers", ctx).Return([]*types.User{user}, nil)
	// call the GetUsers function
	result, err := mock.GetUsers(ctx)
	// assert the result
	assert.Equal(t, []*types.User{user}, result)
	// assert the error
	assert.Nil(t, err)
	// assert that the mock was called
	mock.AssertExpectations(t)

}

// unit test for UpdateUser using mocks from ../mocks/UserStore.go
func TestUpdateUser(t *testing.T) {
	// create a mock object
	mock := mocks.NewUserStore(t)
	// create a new UserStorage object with the mock object
	// create a new context
	ctx := context.Background()

	// mock the UpdateUser function
	mock.On("UpdateUser", ctx, "1", &types.UpdateUserParams{FirstName: "Jane"}).Return(nil)
	// call the UpdateUser function
	err := mock.UpdateUser(ctx, "1", &types.UpdateUserParams{FirstName: "Jane"})
	// assert the error
	assert.Nil(t, err)
	// assert that the mock was called
	mock.AssertExpectations(t)

}

// unit test for DeleteUser using mocks from ../mocks/UserStore.go
func TestDeleteUser(t *testing.T) {
	// create a mock object
	mock := mocks.NewUserStore(t)
	// create a new UserStorage object with the mock object
	// create a new context
	ctx := context.Background()

	// create a new user object
	user := &types.User{
		ID:                "1",
		FirstName:         "John",
		LastName:          "Doe",
		Email:             "john@gmail.com",
		EncryptedPassword: "password",
	}

	// mock the DeleteUser function
	mock.On("DeleteUser", ctx, "1").Return(user, nil)
	// call the DeleteUser function
	result, err := mock.DeleteUser(ctx, "1")
	// assert the result
	assert.Equal(t, user, result)

	// assert the error
	assert.Nil(t, err)
	// assert that the mock was called
	mock.AssertExpectations(t)

}
