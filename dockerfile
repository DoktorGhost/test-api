# Используем официальный образ PostgreSQL
FROM postgres:latest

ENV POSTGRES_USER admin
ENV POSTGRES_PASSWORD admin
ENV POSTGRES_DB admin

# Открытие порта для доступа к PostgreSQL
EXPOSE 5433

# Запуск PostgreSQL при запуске контейнера
CMD ["postgres"]
