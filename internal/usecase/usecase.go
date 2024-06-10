package usecase

import (
	"api/internal/storage"
	"api/internal/validator"
	"errors"
	"github.com/google/uuid"
)

type UseCase struct {
	storage storage.RepositoryDB
}

func NewUseCase(storage storage.RepositoryDB) *UseCase {
	return &UseCase{storage: storage}
}

func (uc *UseCase) UCCreate(user storage.User) (uuid.UUID, error) {
	if err := validator.ValidateStruct(user); err != nil {
		return uuid.Nil, err
	}

	id, err := uc.storage.Create(user)

	return id, err
}

func (uc *UseCase) UCDelete(id uuid.UUID) error {
	err := uc.storage.Delete(id)
	return err
}

func (uc *UseCase) UCRead(id uuid.UUID) (storage.User, error) {
	user, err := uc.storage.Read(id)
	return user, err
}

func (uc *UseCase) UCUpdate(user storage.User) error {
	if user.ID == uuid.Nil {
		return errors.New("invalid ID")
	}

	if err := validator.ValidateStruct(user); err != nil {

		return err
	}

	err := uc.storage.Update(user)
	return err
}

func (uc *UseCase) UCClose() error {
	return uc.storage.Close()
}
