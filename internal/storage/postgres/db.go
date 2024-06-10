package postgres

import (
	"api/internal/storage"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"log"
	"sync"
)

type PostgresStorage struct {
	db *sql.DB
	mu sync.RWMutex
}

// NewPostgresRepository создает новый экземпляр PostgresRepository.
func NewPostgresStorage(dsn string) (*PostgresStorage, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Установка расширения uuid-ossp, если оно не установлено
	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Создание таблицы, если она не существует
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS users (
        ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        Firstname VARCHAR(100) NOT NULL,
        Lastname VARCHAR(100) NOT NULL,
        Email VARCHAR(255) NOT NULL UNIQUE,
        Age INT NOT NULL,
        CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );
    `
	if _, err = db.Exec(createTableQuery); err != nil {
		log.Println(err)
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (r *PostgresStorage) Read(id uuid.UUID) (storage.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var url string
	err := r.db.QueryRow("SELECT url FROM urls WHERE short_url = $1", shortURL).Scan(&url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("url not found")
		}
		return "", err
	}
	return url, nil
}

func (r *PostgresStorage) Create(user storage.User) (string, error) {

}

func (r *PostgresStorage) Delete(shortURL string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := "DELETE FROM urls WHERE short_url = $1 AND url = $1"
	_, err := r.db.Exec(query, shortURL)
	if err != nil {
		return err
	}
	return nil
}
