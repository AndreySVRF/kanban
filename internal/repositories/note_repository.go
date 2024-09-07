package repositories

import "kanban/internal/entities"

// NoteRepository defines methods for interacting with Note entities.
type NoteRepository interface {
	GetNotesByListID(listID int) ([]entities.Note, error)
	GetNoteByID(id int) (entities.Note, error)
	CreateNote(note entities.Note) error
	UpdateNote(note entities.Note) error
	DeleteNote(id int) error
}
