package user

import (
	"context"
	"errors"
	"net/mail"
	"sync"
)

var (
	ErrInvalidName  = errors.New("invalid name: must be between 1 and 30 characters")
	ErrInvalidAge   = errors.New("invalid age: must be between 0 and 150")
	ErrInvalidEmail = errors.New("invalid email format")
)

// User represents a chat user
// TODO: Add more fields if needed

type User struct {
	Name  string
	Email string
	ID    string
}

// Validate checks if the user data is valid, returns an error for each invalid field
func (u *User) Validate() error {
	if !IsValidName(u.Name) {
		return ErrInvalidName
	}

	if !IsValidAge(u.Age) {
		return ErrInvalidAge
	}

	if !IsValidEmail(u.Email) {
		return ErrInvalidEmail
	}

	return nil
}

// IsValidEmail checks if the email format is valid
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// IsValidName checks if the name is valid, returns false if the name is empty or longer than 30 characters
func IsValidName(name string) bool {
	lenName := len(name)
	return 1 <= lenName && lenName <= 30
}

// IsValidAge checks if the age is valid, returns false if the age is not between 0 and 150
func IsValidAge(age int) bool {
	return 0 <= age && age <= 150
}

// UserManager manages users
// Contains a map of users, a mutex, and a context

type UserManager struct {
	ctx   context.Context
	users map[string]User // userID -> User
	mutex sync.RWMutex    // Protects users map
	// TODO: Add more fields if needed
}

// NewUserManager creates a new UserManager
func NewUserManager() *UserManager {
	// TODO: Initialize UserManager fields
	return &UserManager{
		users: make(map[string]User),
	}
}

// NewUserManagerWithContext creates a new UserManager with context
func NewUserManagerWithContext(ctx context.Context) *UserManager {
	// TODO: Initialize UserManager with context
	return &UserManager{
		ctx:   ctx,
		users: make(map[string]User),
	}
}

// AddUser adds a user
func (m *UserManager) AddUser(u User) error {
	// TODO: Add user to map, check context
	return nil
}

// RemoveUser removes a user
func (m *UserManager) RemoveUser(id string) error {
	// TODO: Remove user from map
	return nil
}

// GetUser retrieves a user by id
func (m *UserManager) GetUser(id string) (User, error) {
	// TODO: Get user from map
	return User{}, errors.New("not found")
}
