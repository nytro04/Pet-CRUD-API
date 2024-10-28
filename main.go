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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	var (
		dbUser     = os.Getenv("DB_USER")
		dbName     = os.Getenv("POSTGRES_DB_NAME")
		dbPassword = os.Getenv("POSTGRES_DB_PASSWORD")
		dbPort     = os.Getenv("POSTGRES_DB_PORT")
		host       = os.Getenv("HOST")
		listenAddr = os.Getenv("API_PORT")
	)

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, dbPort, dbUser, dbPassword, dbName,
	)

	dbConn, err := db.NewDB(connStr)
	if err != nil {
		log.Fatalf("Error while connecting to the database: %s\n", err)
	}

	defer dbConn.Close()

	petStore := db.NewPetStorage(dbConn)
	if err := petStore.Init(); err != nil {
		log.Fatalf("Error while initializing pet storage: %s\n", err)
	}

	userStore := db.NewUserStorage(dbConn)
	if err := userStore.Init(); err != nil {
		log.Fatalf("Error while initializing user storage: %s\n", err)
	}

	var (
		store = &db.Store{
			Pet:  petStore,
			User: userStore,
		}

		petHandler  = api.NewPetHandler(store.Pet)
		userHandler = api.NewUserHandler(store.User)
	)

	mux := http.NewServeMux()

	// Pet handlers
	mux.HandleFunc("POST /api/v1/pet", petHandler.CreatePetHandler)
	mux.HandleFunc("GET /api/v1/pet/{id}", petHandler.GetPetByIdHandler)
	mux.HandleFunc("GET /api/v1/pet", petHandler.GetPetsHandler)
	mux.HandleFunc("PATCH /api/v1/pet/{id}", petHandler.UpdatePetHandler)
	mux.HandleFunc("DELETE /api/v1/pet/{id}", petHandler.DeleteHandler)

	// User handlers
	mux.HandleFunc("POST /api/v1/user", userHandler.HandleCreateUser)
	mux.HandleFunc("GET /api/v1/user/{id}", userHandler.HandleGetUser)
	mux.HandleFunc("GET /api/v1/user", userHandler.HandleGetUsers)
	mux.HandleFunc("PATCH /api/v1/user/{id}", userHandler.HandleUpdateUser)
	mux.HandleFunc("DELETE /api/v1/user/{id}", userHandler.HandleDeleteUser)

	log.Printf("Starting server on %s\n", listenAddr)

	err = http.ListenAndServe("localhost:"+listenAddr, mux)
	if err != nil {
		log.Fatalf("Error while starting the server: %s\n", err)
	}

}
