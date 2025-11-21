package cmd

import (
	"database/sql"
	"main/internal/handlers"
	"main/internal/repository/postgres"
	userusecase "main/internal/usecase/user"
)

func main() {
	var db *sql.DB

	userRepo := postgres.NewUserRepository(db)

	userUseCase := userusecase.NewUserUseCase(userRepo)

	userHandler := handlers.NewUserHandler(userUseCase)

	router := http.
}
