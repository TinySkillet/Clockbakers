// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	PhoneNo   string
	Address   string
	Password  string
	Role      string
}
