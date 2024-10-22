package api

import (
	"testing"

	"github.com/nytro04/pet-crud/types"
)

// test Post /api/v1/user/
// func TestHandleCreateUser(t *testing.T) {
// 	tdb := setup(t)
// 	defer tdb.teardown(t)

// 	// test case

// 	userHandler := NewUserHandler(tdb.Store)

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("POST /api/v1/user", userHandler.HandleCreateUser)

// 	params := types.CreateUserParams{
// 		FirstName: "John",
// 		LastName:  "Doe",
// 		Email:     "doe@gmail.com",
// 		Password:  "password",
// 	}

// 	b, _ := json.Marshal(params)
// 	req := httptest.NewRequest("POST", "/api/v1/user", bytes.NewReader(b))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()

// 	var user types.User
// 	json.NewDecoder(w.Body).Decode(&user)

// }

type MockUserService struct {
	HandleCreateUserFunc func(user types.User) (string, error)
	UsersCreated         []types.User
}

func (m *MockUserService) HandleCreateUser(user types.User) (string, error) {
	m.UsersCreated = append(m.UsersCreated, user)
	return m.HandleCreateUserFunc(user)
}

func TestHandleCreateUser(t *testing.T) {
	t.Run("can create user", func(t *testing.T) {
		user := types.User{
			FirstName:         "John",
			LastName:          "Doe",
			Email:             "john@gmail.com",
			EncryptedPassword: "password",
		}
		expectedID := "1"

		mockUserService := &MockUserService{
			HandleCreateUserFunc: func(user types.User) (string, error) {
				return expectedID, nil
			},
		}

		userHandler := NewUserHandler(mockUserService)

	})
}

// func TestHandleCreateUser(t *testing.T) {
// 	// tdb := setup(t)
// 	// defer tdb.teardown(t)

// 	// userHandler := NewUserHandler(tdb.Store)

// 	mux := http.NewServeMux()
// 	// mux.HandleFunc("/api/v1/user", userHandler.HandleCreateUser)

// 	tests := []struct {
// 		name           string
// 		input          types.CreateUserParams
// 		expectedStatus int
// 	}{
// 		{
// 			name: "valid user",
// 			input: types.CreateUserParams{
// 				FirstName: "John",
// 				LastName:  "Doe",
// 				Email:     "doe@gmail.com",
// 				Password:  "password",
// 			},
// 			expectedStatus: http.StatusCreated,
// 		},
// 		{
// 			name: "invalid method",
// 			input: types.CreateUserParams{
// 				FirstName: "Jane",
// 				LastName:  "Doe",
// 				Email:     "jane@gmail.com",
// 				Password:  "password",
// 			},
// 			expectedStatus: http.StatusMethodNotAllowed,
// 		},
// 		{
// 			name: "invalid input",
// 			input: types.CreateUserParams{
// 				FirstName: "",
// 				LastName:  "Doe",
// 				Email:     "invalid-email",
// 				Password:  "password",
// 			},
// 			expectedStatus: http.StatusBadRequest,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			b, _ := json.Marshal(tt.input)
// 			req := httptest.NewRequest("POST", "/api/v1/user", bytes.NewReader(b))
// 			req.Header.Set("Content-Type", "application/json")
// 			w := httptest.NewRecorder()

// 			mux.ServeHTTP(w, req)

// 			if w.Code != tt.expectedStatus {
// 				t.Errorf("expected status %v, got %v", tt.expectedStatus, w.Code)
// 			}

// 			if tt.expectedStatus == http.StatusCreated {
// 				var user types.User
// 				if err := json.NewDecoder(w.Body).Decode(&user); err != nil {
// 					t.Errorf("failed to decode response: %v", err)
// 				}
// 				if user.Email != tt.input.Email {
// 					t.Errorf("expected email %v, got %v", tt.input.Email, user.Email)
// 				}
// 			}
// 		})
// 	}
// }
