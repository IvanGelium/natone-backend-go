package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/IvanGelium/main-service/internal/handlers"
	"github.com/IvanGelium/main-service/internal/middleware"
	"github.com/IvanGelium/main-service/internal/repositories/postgress"
	"github.com/IvanGelium/main-service/internal/services"
	pgxpool "github.com/jackc/pgx/v5/pgxpool"
	goose "github.com/pressly/goose/v3"
)

func main() {
	dsn := "postgres://postgres:password@localhost:5432/my_db?sslmode=disable"

	migrationDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения для миграций: %v", err)
	}
	defer migrationDB.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Ошибка настройки диалекта goose: %v", err)
	}

	log.Println("Запуск миграций базы данных...")
	if err := goose.Up(migrationDB, "db/migrations"); err != nil {
		log.Fatalf("Ошибка применения миграций: %v", err)
	}
	log.Println("Миграции успешно применены!")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbPool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Не удалось создать пул соединений: %v", err)
	}
	defer dbPool.Close()

	if err := dbPool.Ping(ctx); err != nil {
		log.Fatalf("База данных недоступна: %v", err)
	}
	log.Println("Успешно подключились к PostgreSQL через pgxpool!")

	authRepo := postgress.NewAuthRepository(dbPool)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /auth/signup", authHandler.SignUp)

	// Пример защищенного роута (оборачиваем его в нашу мидлвару)
	secretSecret := []byte("secret_key")
	protectedMux := middleware.AuthMiddleware(secretSecret)(mux)

	log.Println("Бэкенд по Чистой Архитектуре запущен на http://localhost:5001")
	if err := http.ListenAndServe(":5001", protectedMux); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}
