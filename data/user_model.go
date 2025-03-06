package models

import (
	"regexp"

	"github.com/TinySkillet/ClockBakers/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	PhoneNo   string    `json:"phone_no" validate:"required,phone"`
	Address   string    `json:"address" validate:"required"`
	Password  string    `json:"password,omitempty" validate:"required,password"`
	Role      string    `json:"role" validate:"required,role"`
}

func (u *User) Validate() {
	var Validate = validator.New()
	Validate.RegisterValidation("email", validateEmail)
	Validate.RegisterValidation("phone", validatePhone)
	Validate.RegisterValidation("password", validatePassword)
	Validate.RegisterValidation("role", validateRole)
}

func validateEmail(f1 validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(f1.Field().String())
}

func validatePhone(f1 validator.FieldLevel) bool {
	re := regexp.MustCompile(`^\+?[1-9]\d{0,2}[-.\s]?\(?\d{1,4}\)?[-.\s]?\d{1,4}[-.\s]?\d{1,9}$`)
	return re.MatchString(f1.Field().String())
}

func validateRole(f1 validator.FieldLevel) bool {
	roles := []string{"admin", "customer"}

	for _, role := range roles {
		if role == f1.Field().String() {
			return true
		}
	}
	return false
}

func validatePassword(f1 validator.FieldLevel) bool {
	// at least 1 uppercase, 1 lowercase, 1 digit, 1 special and at least 8 characters long
	re := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()\-_=+]).{8,}$`)
	return re.MatchString(f1.Field().String())
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func DBUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
		PhoneNo:   dbUser.PhoneNo,
		Address:   dbUser.Address,
		Password:  "",
		Role:      dbUser.Role,
	}
}

func DBUsersToUsers(dbUsers []database.User) []User {
	users := make([]User, len(dbUsers))
	for i, dbUser := range dbUsers {
		users[i] = DBUserToUser(dbUser)
	}
	return users
}
