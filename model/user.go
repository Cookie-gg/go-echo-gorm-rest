package model

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" param:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var UserColumns = []string{
	"id",
	"name",
	"created_at",
	"updated_at",
}
