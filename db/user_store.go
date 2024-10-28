package db

import (
	"context"
	"database/sql"

	"github.com/nytro04/pet-crud/types"
)

type UserStorage struct {
	db *sql.DB
}

// constructor/factory function
func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (s *UserStorage) Init() error {
	return s.createUserTable()
}

func (s *UserStorage) createUserTable() error {
	query := `create table if not exists users (
		id serial primary key,
		FirstName varchar(100),
		LastName varchar(100),
		Email varchar(100),
		Password varchar(100)
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *UserStorage) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {

	query := `INSERT INTO users
		(firstName, lastName, email, password)
		values ($1, $2, $3, $4)
		RETURNING id
	`
	err := s.db.QueryRowContext(ctx, query, user.FirstName, user.LastName, user.Email, user.EncryptedPassword).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserStorage) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	query := `SELECT id, firstName, lastName, email, password
		FROM users
		WHERE email = $1
	`
	var user types.User
	err := s.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.EncryptedPassword)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserStorage) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	query := `SELECT id, firstName, lastName, email, password
		FROM users
		WHERE id = $1
	`
	var user types.User

	err := s.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.EncryptedPassword)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserStorage) GetUsers(ctx context.Context) ([]*types.User, error) {
	query := `SELECT id, firstName, lastName, email, password
		FROM users
	`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*types.User
	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.EncryptedPassword)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return users, nil
}

func (s *UserStorage) UpdateUser(ctx context.Context, id string, params *types.UpdateUserParams) error {
	query := `UPDATE users
		SET firstName = $1, lastName = $2
		WHERE id = $3
	`
	_, err := s.db.ExecContext(ctx, query, params.FirstName, params.LastName, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStorage) DeleteUser(ctx context.Context, id string) (*types.User, error) {
	query := `DELETE FROM users
		WHERE id = $1
	`

	var user types.User

	err := s.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.EncryptedPassword)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
