<p align="left">
      <img src="https://i.ibb.co/cYzQsPG/logoza-ru.png" alt="Project Logo" width="726">
</p>

## Настройка переменных окружения
Если запускаем через контейнер, то .env не меняем.

В противном случае:

```golang
DB_HOST: localhost #хост для бд
DB_PORT: "5432" #порт для бд
DB_USER: admin #имя пользователя бд
DB_PASSWORD: admin #пароль пользователя бд
DB_NAME: admin #имя бд
API_PORT: "8080" #порт сервера
API_HOST: localhost #хост сервера
```
## Запуск контейнера
Собираем образ и поднимаем контейнер:

```golang
docker build -t img .
docker run -d --name pg-con -p 5432:5432 img
```

## Запуск приложения
перходим в CMD/server и запускаем приложение 
```golang
go run main.go
```

## Тестирование
Протестировать можно в Postman

1. тип POST , http://localhost:8080/users , в теле отправляем  JSON вида:

  ```JSON
  {
	  "firstname": "Oleg",
	  "lastname" : "Samsonovs",
	  "email": "ghostkot@egmail.com",
	  "age": 150
  }
  ```
2. тип PUT , localhost:8080/users/ , в теле запроса отправляем JSON с измененными данными, включая ID, например:

```JSON
{
    "id": "387c955c-a955-4a1c-8767-3c1b81049e7b",
    "firstname": "Oleячсgator",
    "lastname": "Saячсячсmsonovsssss",
    "email": "ghot@gmail.com",
    "age": 14,
    "created": "2024-06-10T19:38:37.954632Z"
}
```
3. Тип GET , localhost:8080/users/{id}  ,  id - записи, которую мы хотим получить. в отает получаем JSON
4. Тип DELETE, localhost:8080/users/{id} , удаление записи по id
