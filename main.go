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
	connStr := "user=postgres dbname=postgres password=petDBsecr3t sslmode=disable"

	dbConn, err := db.NewDB(connStr)
	if err != nil {
		log.Fatalf("Error while connecting to the database: %s\n", err)

	}

	defer dbConn.Close()
	fmt.Println("database ok!!")

	// petStore, err := db.NewPostgresPetStore(dbConn)
	// petStore.
	// if err != nil {
	// }

	petStore := db.NewPetStorage(dbConn)
	petStore.Init()

	var (
		store = &db.Store{
			Pet: petStore,
		}

		petHandler = api.NewPetHandler(store)

		app   = fiber.New()
		apiV1 = app.Group("/api/v1")
	)

	// Pet handlers
	apiV1.Post("/pet", petHandler.CreatePetHandler)
	apiV1.Get("/pet/:id", petHandler.GetPetByIdHandler)
	apiV1.Get("/pet", petHandler.GetPetsHandler)
	apiV1.Patch("/pet/:id", petHandler.UpdatePetHandler)
	apiV1.Delete("/pet/:id", petHandler.DeleteHandler)

	listenAddr := os.Getenv("API_PORT")

	app.Listen(listenAddr)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
