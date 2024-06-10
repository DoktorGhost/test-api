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
