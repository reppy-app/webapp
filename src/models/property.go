package models

import "time"

// Property represents a table property on database.
type Property struct {
	ID        uint64    `json:"id,omitempty"`
	UserID    uint64    `json:"userId,omitempty"`
	Name      string    `json:"name,omitempty"`
	Price     float64   `json:"price,omitempty"`
	CreatedAt time.Time `json:"createAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	UserEmail string    `json:"userEmail,omitempty"`
}
