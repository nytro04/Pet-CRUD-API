package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nytro04/pet-crud/api"
	"github.com/nytro04/pet-crud/db"
)

func main() {

	var (
		dbUser     = os.Getenv("DB_USER")
		dbName     = os.Getenv("POSTGRES_DB_NAME")
		dbPassword = os.Getenv("POSTGRES_DB_PASSWORD")
	)

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("failed to load env", err)
	// }

	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPassword,
	)

	dbConn, err := db.NewDB(connStr)
	if err != nil {
		log.Fatalf("Error while connecting to the database: %s\n", err)

	}

	defer dbConn.Close()

	petStore := db.NewPetStorage(dbConn)
	petStore.Init()

	userStore := db.NewUserStorage(dbConn)
	userStore.Init()

	var (
		store = &db.Store{
			Pet:  petStore,
			User: userStore,
		}

		petHandler = api.NewPetHandler(store)
		// userHandler = api.NewUserHandler(store)

		// app   = fiber.New()
		// apiV1 = app.Group("/api/v1")

	)

	mux := http.NewServeMux()
	// Pet handlers
	mux.HandleFunc("POST /api/v1/pet/", petHandler.CreatePetHandler)
	mux.HandleFunc("GET /api/v1/pet/{id}", petHandler.GetPetByIdHandler)
	mux.HandleFunc("GET /api/v1/pet/", petHandler.GetPetsHandler)
	mux.HandleFunc("PATCH /api/v1/pet/{id}", petHandler.UpdatePetHandler)
	mux.HandleFunc("DELETE /api/v1/pet/{id}", petHandler.DeleteHandler)

	// // User handlers
	// apiV1.Post("/user", userHandler.HandleCreateUser)
	// apiV1.Get("/user/:id", userHandler.HandleGetUser)
	// apiV1.Get("/user", userHandler.HandleGetUsers)
	// apiV1.Patch("/user/:id", userHandler.HandleUpdateUser)
	// apiV1.Delete("/user/:id", userHandler.HandleDeleteUser)

	listenAddr := os.Getenv("API_PORT")

	log.Printf("Starting server on %s\n", listenAddr)

	err = http.ListenAndServe(listenAddr, mux)
	if err != nil {
		log.Fatalf("Error while starting the server: %s\n", err)
	}

	// app.Listen(listenAddr)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
