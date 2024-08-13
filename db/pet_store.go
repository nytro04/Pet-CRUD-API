package db

import (
	"context"
	"database/sql"

	"github.com/nytro04/pet-crud/types"
)

type PetStore interface {
	CreatePet(context.Context, *types.Pet) (*types.Pet, error)
	Close() error
}

type PetStorage struct {
	db *sql.DB
}

func NewPetStorage(db *sql.DB) *PetStorage {
	return &PetStorage{
		db: db,
	}
}

func (s *PetStorage) Init() error {
	return s.createPetTable()
}

func (db *PetStorage) Close() error {
	return db.db.Close()
}

func (s *PetStorage) createPetTable() error {
	query := `create table if not exists pet (
		id serial primary key,
		name varchar(100),
		owner varchar(100),
		type varchar(100),
		age serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PetStorage) CreatePet(ctx context.Context, pet *types.Pet) (*types.Pet, error) {
	query := `INSERT INTO pet
		(name, owner, type, age, created_at)
		values ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := s.db.QueryRowContext(ctx, query, pet.Name, pet.Owner, pet.Type, pet.Age, pet.CreatedAt).
		Scan(&pet.ID)
	if err != nil {
		return nil, err
	}

	return pet, nil
}
