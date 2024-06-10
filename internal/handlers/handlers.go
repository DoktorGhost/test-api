package handlers

import (
	"api/internal/storage"
	"api/internal/usecase"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func InitRoutes(useCase usecase.UseCase) chi.Router {
	r := chi.NewRouter()

	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		handlerCreate(w, r, useCase)
	})
	r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlerRead(w, r, useCase)
	})
	r.Put("/users/", func(w http.ResponseWriter, r *http.Request) {
		handlerUpdate(w, r, useCase)
	})
	r.Delete("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlerDelete(w, r, useCase)
	})
	return r
}

func handlerCreate(w http.ResponseWriter, r *http.Request, usecase usecase.UseCase) {
	if r.Method != http.MethodPost {
		log.Println("Ошибка метода")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user storage.User
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&user); err != nil {
		log.Println("Ошибка декодирования JSON")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	id, err := usecase.UCCreate(user)
	if err != nil {
		log.Println("Ошибка отправки формы: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stringID := fmt.Sprintf("ID: %v", id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(stringID))

}

func handlerRead(w http.ResponseWriter, r *http.Request, usecase usecase.UseCase) {
	if r.Method != http.MethodGet {
		log.Println("Ошибка метода")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := chi.URLParam(r, "id")

	if id == "" {
		log.Println("Пустой ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		log.Println("Ошибка парсинга ID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := usecase.UCRead(parsedID)
	if err != nil {
		log.Println("Ошибка чтения записи с ID: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Println("Ошибка кодирования JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

func handlerUpdate(w http.ResponseWriter, r *http.Request, usecase usecase.UseCase) {
	if r.Method != http.MethodPut {
		log.Println("Ошибка метода")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user storage.User
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&user); err != nil {
		log.Println("Ошибка декодирования JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err := usecase.UCUpdate(user)
	if err != nil {
		log.Println("Ошибка обновления записи: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handlerDelete(w http.ResponseWriter, r *http.Request, usecase usecase.UseCase) {
	if r.Method != http.MethodDelete {
		log.Println("Ошибка метода")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := chi.URLParam(r, "id")

	if id == "" {
		log.Println("Пустой ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		log.Println("Ошибка парсинга ID: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = usecase.UCDelete(parsedID)
	if err != nil {
		log.Println("Ошибка удаления: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
