package repositories

import "kanban/internal/entities"

// UserRepository defines methods for interacting with User entities.
type UserRepository interface {
	GetUserByUsername(username string) (entities.User, error)
	CreateUser(user entities.User) error
}
