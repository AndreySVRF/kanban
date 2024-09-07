package repositories

import "kanban/internal/entities"

// BoardRepository defines methods for interacting with Board entities.
type BoardRepository interface {
	GetBoards() ([]entities.Board, error)
	GetBoardByID(id int) (entities.Board, error)
	CreateBoard(board entities.Board) error
	UpdateBoard(board entities.Board) error
	DeleteBoard(id int) error
}
