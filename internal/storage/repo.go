package storage

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	Created   time.Time `json:"created"`
}

// Repository представляет интерфейс для работы с хранилищем данных.
type RepositoryDB interface {
	Create(user User) (uuid.UUID, error)
	Read(id uuid.UUID) (User, error)
	Update(user User) error
	Delete(id uuid.UUID) error
	Close() error
}
