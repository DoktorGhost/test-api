package usecase

import (
	"api/internal/mocks"
	"api/internal/storage"
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

	testUser_1 := storage.User{
		ID:        id,
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john@example.com",
		Age:       30,
		Created:   time.Now(),
	}

	testUser_2 := storage.User{
		ID:        id,
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "johnexample.com",
		Age:       30,
		Created:   time.Now(),
	}

	// Указываем ожидания тест 1
	m.EXPECT().Create(testUser_1).Return(id, nil)

	// Вызываем тестируемый метод
	uuid, err := uc.UCCreate(testUser_1)
	require.NoError(t, err)
	require.Equal(t, uuid, id)

	m.EXPECT().Create(testUser_2).Return(id, err)

	_, err2 := uc.UCCreate(testUser_2)
	require.Equal(t, err, err2)
}
