package api

import (
	"database/sql"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/nytro04/pet-crud/db"
)

type testDB struct {
	db *sql.DB
	*db.Store
}

var (
	dbUser     = os.Getenv("DB_USER")
	dbName     = os.Getenv("POSTGRES_TEST_DB_NAME")
	dbPassword = os.Getenv("POSTGRES_DB_PASSWORD")
)

func (tdb *testDB) teardown(t *testing.T) {
	// drop the table
	_, err := tdb.db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		t.Fatalf("Error dropping table: %s\n", err)
	}
}

// test function
func setup(t *testing.T) *testDB {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatal("failed to load env", err)
	}

	// connStr := fmt.Sprintf(
	// 	"user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPassword,
	// )

	connStr := "host=localhost port=5432 user=postgres password=petDBsecr3t dbname=test_postgres sslmode=disable"

	dbConn, err := db.NewDB(connStr)
	if err != nil {
		t.Fatalf("Error while connecting to the database: %s\n", err)
	}

	petStore := db.NewPetStorage(dbConn)
	petStore.Init()

	userStore := db.NewUserStorage(dbConn)
	userStore.Init()

	store := &db.Store{
		Pet:  petStore,
		User: userStore,
	}

	tdb := &testDB{
		db:    dbConn,
		Store: store,
	}

	return tdb
}
