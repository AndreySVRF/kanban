package usecases

import (
	"kanban/internal/entities"
	"kanban/internal/repositories"
)

type BoardUseСase interface {
	GetBoards() ([]entities.Board, error)
	GetBoardByID(id int) (entities.Board, error)
	CreateBoard(title string) error
	UpdateBoard(id int, title string) error
	DeleteBoard(id int) error
}

type boardUseСase struct {
	boardRepo repositories.BoardRepository
}

func NewBoardUseСase(boardRepo repositories.BoardRepository) BoardUseСase {
	return &boardUseСase{
		boardRepo: boardRepo,
	}
}

func (b *boardUseСase) GetBoards() ([]entities.Board, error) {
	return b.boardRepo.GetBoards()
}

func (b *boardUseСase) GetBoardByID(id int) (entities.Board, error) {
	return b.boardRepo.GetBoardByID(id)
}

func (b *boardUseСase) CreateBoard(title string) error {
	board := entities.Board{
		Title: title,
	}
	return b.boardRepo.CreateBoard(board)
}

func (b *boardUseСase) UpdateBoard(id int, title string) error {
	board := entities.Board{
		ID:    id,
		Title: title,
	}
	return b.boardRepo.UpdateBoard(board)
}

func (b *boardUseСase) DeleteBoard(id int) error {
	return b.boardRepo.DeleteBoard(id)
}
