package cmd

import (
	"log"
	"main/internal/config"
	"main/internal/database"
	"main/internal/delivery"
	"main/internal/handlers"
	"main/internal/repository/gorm"
	prusecase "main/internal/usecase/pullrequest"
	teamusecase "main/internal/usecase/team"
	userusecase "main/internal/usecase/user"
)

func main() {
	// загрузка конфига
	cfg := config.Load()

	// Подключение БД
	db, err := database.NewPostgresConn(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("ошибка при подключении к БД: ", err)
	}

	// инициализация репозиториев
	userRepo := gorm.NewUserRepository(db)
	teamRepo := gorm.NewTeamRepository(db)
	prRepo := gorm.NewPRRepository(db)

	// инициализация usecase слоя
	userUseCase := userusecase.NewUserUseCase(userRepo)
	teamUseCase := teamusecase.NewTeamUseCase(teamRepo, userRepo)
	prUseCase := prusecase.NewPullRequestUseCase(prRepo, userRepo, teamRepo)
	_, _ = teamUseCase, prUseCase

	// иницилазиация хендлеров
	userHandler := handlers.NewUserHandler(userUseCase)

	// инициализация роутера
	router := delivery.NewRouter(userHandler)
	ginRouter := router.SetupRoutes()

	// запуск сервера
	log.Printf("запуск сервера на порту %s", cfg.Port)
	ginRouter.Run(":" + cfg.Port)
}
