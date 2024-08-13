package types

import (
	"time"
)

type Pet struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	Type      string    `json:"type"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"createdAt"`
}
type CreatePetParams struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Type  string `json:"type"`
	Age   int    `json:"age"`
}
