package db

import (
	"context"
	"database/sql"

	"github.com/nytro04/pet-crud/types"
)

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

func (s *PetStorage) GetPetById(ctx context.Context, id string) (*types.Pet, error) {
	query := `
		SELECT id, name, owner, type, age, created_at
		FROM pet
		WHERE id = $1
	`

	var pet types.Pet

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&pet.ID,
		&pet.Name,
		&pet.Owner,
		&pet.Type,
		&pet.Age,
		&pet.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &pet, nil
}

func (s *PetStorage) GetPets(ctx context.Context) ([]*types.Pet, error) {
	query := `SELECT id, name, owner, type, age, created_at
						FROM pet
`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pets []*types.Pet
	for rows.Next() {
		var pet types.Pet
		err := rows.Scan(&pet.ID, &pet.Name, &pet.Owner, &pet.Type, &pet.Age, &pet.CreatedAt)
		if err != nil {
			return nil, err
		}
		pets = append(pets, &pet)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return pets, nil
}

func (s *PetStorage) UpdatePet(ctx context.Context, id string, pet *types.CreatePetParams) error {
	query := `UPDATE pet
		SET name = $1, owner = $2, type = $3, age = $4
		WHERE id = $5
	`
	_, err := s.db.ExecContext(ctx, query, pet.Name, pet.Owner, pet.Type, pet.Age, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PetStorage) DeletePet(ctx context.Context, id string) (*types.Pet, error) {
	query := `DELETE FROM pet WHERE id = $1
						RETURNING id, name, owner, type, age, created_at
						`
	var pet types.Pet

	err := s.db.QueryRowContext(ctx, query, id).
		Scan(&pet.ID, &pet.Name, &pet.Owner, &pet.Type, &pet.Age, &pet.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &pet, nil

}
