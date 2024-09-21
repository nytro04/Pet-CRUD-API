package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nytro04/pet-crud/types"
)

// test Post /api/v1/user/
func TestHandleCreateUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	// test case

	userHandler := NewUserHandler(tdb.Store)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/user", userHandler.HandleCreateUser)

	params := types.CreateUserParams{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "doe@gmail.com",
		Password:  "password",
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/api/v1/user/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	var user types.User
	json.NewDecoder(w.Body).Decode(&user)

}
