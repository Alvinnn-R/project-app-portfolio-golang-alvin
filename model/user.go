package model

import "time"

// User represents an admin user
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"` // Password hash, never exposed in JSON
	Name      string    `json:"name"`
	Role      string    `json:"role"` // admin, user
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
