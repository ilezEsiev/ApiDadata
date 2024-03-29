# Используем образ Golang для сборки приложения
FROM golang:latest

# Устанавливаем переменную окружения для задания рабочей директории внутри контейнера
WORKDIR /app

# Копируем все файлы из текущего каталога внутрь контейнера
COPY . .

# Собираем приложение
RUN go build -o main .

# Указываем, что приложение слушает порт 8080
EXPOSE 8080

# Команда для запуска приложения при старте контейнера
CMD ["./main"]