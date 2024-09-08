package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

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

		petHandler  = api.NewPetHandler(store)
		userHandler = api.NewUserHandler(store)

		app   = fiber.New()
		apiV1 = app.Group("/api/v1")
	)

	// Pet handlers
	apiV1.Post("/pet", petHandler.CreatePetHandler)
	apiV1.Get("/pet/:id", petHandler.GetPetByIdHandler)
	apiV1.Get("/pet", petHandler.GetPetsHandler)
	apiV1.Patch("/pet/:id", petHandler.UpdatePetHandler)
	apiV1.Delete("/pet/:id", petHandler.DeleteHandler)

	// User handlers
	apiV1.Post("/user", userHandler.HandleCreateUser)
	apiV1.Get("/user/:id", userHandler.HandleGetUser)
	apiV1.Get("/user", userHandler.HandleGetUsers)
	apiV1.Patch("/user/:id", userHandler.HandleUpdateUser)
	apiV1.Delete("/user/:id", userHandler.HandleDeleteUser)

	listenAddr := os.Getenv("API_PORT")

	app.Listen(listenAddr)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
