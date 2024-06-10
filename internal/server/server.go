package server

import (
	"api/internal/config"
	"api/internal/handlers"
	"api/internal/storage/postgres"
	"api/internal/usecase"
	"log"
	"net/http"
)

func StartServer(conf *config.Config) error {
	db, err := postgres.NewPostgresStorage(conf.DatabaseDSN)
	if err != nil {
		log.Fatal("Ошибка подключения к БД: ", err)
	}
	log.Println("Успешное подключение к БД")
	UseCaseDB := usecase.NewUseCase(db)

	r := handlers.InitRoutes(*UseCaseDB)

	//запускаем сервер
	err = http.ListenAndServe(conf.Host+":"+conf.Port, r)
	if err != nil {
		log.Fatal("Ошибка старта сервера: ", err)
	}

	return nil
}
