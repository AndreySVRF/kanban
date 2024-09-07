package repositories

import "kanban/internal/entities"

// ListRepository defines methods for interacting with List entities.
type ListRepository interface {
	GetListsByBoardID(boardID int) ([]entities.List, error)
	GetListByID(id int) (entities.List, error)
	CreateList(list entities.List) error
	UpdateList(list entities.List) error
	DeleteList(id int) error
}
