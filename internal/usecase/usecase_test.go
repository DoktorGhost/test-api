package usecase

import (
	"api/internal/mocks"
	"api/internal/storage"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUCCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRepositoryDB(ctrl)

	uc := NewUseCase(m)

	id := uuid.New()

	tests := []struct {
		name      string
		user      storage.User
		mockCalls int
		mockErr   error
	}{
		{
			name:      "Valid user",
			user:      storage.User{ID: id, Firstname: "John", Lastname: "Doe", Email: "john@example.com", Age: 30, Created: time.Now()},
			mockCalls: 1,
			mockErr:   nil,
		},
		{
			name:      "Invalid email",
			user:      storage.User{ID: id, Firstname: "John", Lastname: "Doe", Email: "johnexample.ru", Age: 30, Created: time.Now()},
			mockCalls: 0,
			mockErr:   errors.New("invalid email"),
		},
		{
			name:      "Invalid name",
			user:      storage.User{ID: id, Firstname: "John2", Lastname: "Doe", Email: "johne@xample.ru", Age: 30, Created: time.Now()},
			mockCalls: 0,
			mockErr:   errors.New("invalid name"),
		},
		{
			name:      "Invalid lastname",
			user:      storage.User{ID: id, Firstname: "John", Lastname: "Doe22", Email: "johne@xample.ru", Age: 30, Created: time.Now()},
			mockCalls: 0,
			mockErr:   errors.New("invalid lastname"),
		},
		{
			name:      "Invalid age",
			user:      storage.User{ID: id, Firstname: "John", Lastname: "Doe", Email: "johne@xample.ru", Age: 152, Created: time.Now()},
			mockCalls: 0,
			mockErr:   errors.New("invalid age"),
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.EXPECT().Create(tt.user).Times(tt.mockCalls).Return(id, tt.mockErr)

			uuidd, err := uc.UCCreate(tt.user)

			if i == 0 {
				require.NoError(t, err)
				require.Equal(t, uuidd, id)
			} else if i == 1 {
				require.Error(t, err)
				require.Equal(t, "invalid email", err.Error())
				require.Equal(t, uuidd, uuid.Nil)
			} else if i == 2 {
				require.Error(t, err)
				require.Equal(t, "invalid name", err.Error())
				require.Equal(t, uuidd, uuid.Nil)
			} else if i == 3 {
				require.Error(t, err)
				require.Equal(t, "invalid lastname", err.Error())
				require.Equal(t, uuidd, uuid.Nil)
			} else if i == 4 {
				require.Error(t, err)
				require.Equal(t, "invalid age", err.Error())
				require.Equal(t, uuidd, uuid.Nil)
			}
		})
	}
}

func TestUCDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRepositoryDB(ctrl)
	uc := NewUseCase(m)

	// Тест на успешное удаление
	t.Run("Successful Deletion", func(t *testing.T) {
		testID := uuid.New()
		m.EXPECT().Delete(testID).Return(nil)

		err := uc.UCDelete(testID)
		require.NoError(t, err)
	})

	// Тест на ошибочное удаление
	t.Run("Deletion Error", func(t *testing.T) {
		testID := uuid.New()
		expectedErr := errors.New("")
		m.EXPECT().Delete(testID).Return(expectedErr)

		err := uc.UCDelete(testID)
		require.Error(t, err)
		require.Equal(t, expectedErr, err)
	})
}

func TestUCRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRepositoryDB(ctrl)
	uc := NewUseCase(m)

	id := uuid.New()
	expectedUser := storage.User{
		ID:        id,
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john@example.com",
		Age:       30,
		Created:   time.Now(),
	}

	// Тест на успешное чтение
	t.Run("Successful Read", func(t *testing.T) {
		m.EXPECT().Read(id).Return(expectedUser, nil)

		user, err := uc.UCRead(id)
		require.NoError(t, err)
		require.Equal(t, expectedUser, user)
	})

	// Тест на ошибочное чтение
	t.Run("Read Error", func(t *testing.T) {
		expectedErr := errors.New("read error")
		m.EXPECT().Read(id).Return(storage.User{}, expectedErr)

		user, err := uc.UCRead(id)
		require.Error(t, err)
		require.Equal(t, expectedErr, err)
		require.Equal(t, storage.User{}, user)
	})
}

func TestUCUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRepositoryDB(ctrl)
	uc := NewUseCase(m)

	validUser := storage.User{
		ID:        uuid.New(),
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john@example.com",
		Age:       30,
		Created:   time.Now(),
	}

	invalidUser := storage.User{
		ID:        uuid.Nil,
		Firstname: "Jane",
		Lastname:  "Doe",
		Email:     "jane@example.com",
		Age:       25,
		Created:   time.Now(),
	}

	invalidEmailUser := storage.User{
		ID:        uuid.New(),
		Firstname: "Jake",
		Lastname:  "Doe",
		Email:     "invalid-email",
		Age:       35,
		Created:   time.Now(),
	}

	// Тест на успешное обновление
	t.Run("Successful Update", func(t *testing.T) {
		m.EXPECT().Update(validUser).Return(nil)

		err := uc.UCUpdate(validUser)
		require.NoError(t, err)
	})

	// Тест на обновление с неверным ID
	t.Run("Update with Invalid ID", func(t *testing.T) {
		err := uc.UCUpdate(invalidUser)
		require.Error(t, err)
		require.Equal(t, err.Error(), "invalid ID")
	})

	// Тест на обновление с неверным email
	t.Run("Update with Invalid Email", func(t *testing.T) {
		err := uc.UCUpdate(invalidEmailUser)
		require.Error(t, err)
		require.Equal(t, err.Error(), "invalid email")
	})

	// Тест на ошибку обновления в базе данных
	t.Run("Update Database Error", func(t *testing.T) {
		expectedErr := errors.New("update error")
		m.EXPECT().Update(validUser).Return(expectedErr)

		err := uc.UCUpdate(validUser)
		require.Error(t, err)
		require.Equal(t, expectedErr, err)
	})
}
