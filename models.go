// (OPTIONAL - You can avoid this file) - Personal preference
// This file makes SQLC types to snake-case types when responded in JSON.
package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/preetDev004/rss-aggregator/db"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func dbUserToUser(user db.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}