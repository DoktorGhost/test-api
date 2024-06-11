package postgres

import (
	"api/internal/config"
	"api/internal/storage"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"time"
)

type PostgresStorage struct {
	db *sql.DB
	mu sync.RWMutex
}

// NewPostgresRepository создает новый экземпляр PostgresRepository.
func NewPostgresStorage(conf *config.Config) (*PostgresStorage, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPass, conf.DBName)
	db, err := sql.Open("postgres", psqlInfo)
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
        Firstname VARCHAR(50) NOT NULL,
        Lastname VARCHAR(50) NOT NULL,
        Email VARCHAR(255) NOT NULL UNIQUE,
        Age INT NOT NULL,
        Created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );
    `
	if _, err = db.Exec(createTableQuery); err != nil {
		log.Println("Ошибка создания таблицы: ", err)
		return nil, err
	}
	log.Println("Таблица успешно создана")
	return &PostgresStorage{db: db}, nil
}

func (r *PostgresStorage) Create(user storage.User) (uuid.UUID, error) {
	//генерируем ID
	user.ID = uuid.New()
	//время создания записи
	user.Created = time.Now()

	query := `
    INSERT INTO users (id, firstname, lastname, email, age, created)
    VALUES ($1, $2, $3, $4, $5, $6)`

	//проверка ID на уникальность в БД
	for {
		_, err := r.db.Exec(query, user.ID, user.Firstname, user.Lastname, user.Email, user.Age, user.Created)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
				// код ошибки уникальности "23505" в Postgres, генерируем по новой
				user.ID = uuid.New()
				continue
			}
			return uuid.Nil, err
		}
		break
	}
	log.Println("Запись добавлена. ID: ", user.ID)
	return user.ID, nil
}

func (r *PostgresStorage) Read(id uuid.UUID) (storage.User, error) {
	var user storage.User
	r.mu.RLock()
	defer r.mu.RUnlock()

	err := r.db.QueryRow("SELECT id, firstname, lastname, email, age, created FROM users WHERE ID = $1", id).Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Age, &user.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.User{}, fmt.Errorf("user with ID %s not found", id)
		}
		return storage.User{}, err
	}
	return user, nil
}

func (r *PostgresStorage) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Ошибка удаления записи с ID: ", id)
		return err
	}
	log.Println("Запсь удалена. ID: ", id)
	return nil
}

func (r *PostgresStorage) Update(user storage.User) error {
	r.mu.Lock() // Locking for writing
	defer r.mu.Unlock()

	query := `
    UPDATE users
    SET firstname = $2, lastname = $3, email = $4, age = $5
    WHERE id = $1`

	_, err := r.db.Exec(query, user.ID, user.Firstname, user.Lastname, user.Email, user.Age)
	if err != nil {
		log.Println("Ошибка обновления записи с ID: ", user.ID)
		return err
	}
	log.Println("Запсь обновлена. ID: ", user.ID)
	return nil
}

func (ps *PostgresStorage) Close() error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	return ps.db.Close()
}
