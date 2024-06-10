package storage

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID
	Firstname string
	Lastname  string
	Email     string
	Age       uint
	Created   time.Time
}

// Repository представляет интерфейс для работы с хранилищем данных.
type RepositoryDB interface {
	Create(user User) (string, error)
	Read(id uuid.UUID) (User, error)
	Put(user User) error
	Delete(id uuid.UUID) error
}
