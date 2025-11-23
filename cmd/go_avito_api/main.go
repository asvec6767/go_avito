package main

import (
	"log"
	"main/internal/config"
	"main/internal/database"
	"main/internal/delivery"
	"main/internal/delivery/handlers"
	"main/internal/repository/gorm"
	prusecase "main/internal/usecase/pullrequest"
	teamusecase "main/internal/usecase/team"
	userusecase "main/internal/usecase/user"

	"github.com/gin-gonic/gin"
)

func main() {
	// загрузка конфига
	cfg := config.Load()

	// настройка типа логирования Gin
	if cfg.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Подключение БД
	db, err := database.NewPostgresConn(cfg.GetDataBaseURL())
	if err != nil {
		log.Fatal("ошибка при подключении к БД: ("+cfg.GetDataBaseURL()+")", err)
	}

	// инициализация репозиториев
	userRepo := gorm.NewUserRepository(db)
	teamRepo := gorm.NewTeamRepository(db)
	prRepo := gorm.NewPRRepository(db)

	// инициализация usecase слоя
	userUseCase := userusecase.NewUserUseCase(userRepo, teamRepo)
	teamUseCase := teamusecase.NewTeamUseCase(teamRepo, userRepo)
	prUseCase := prusecase.NewPullRequestUseCase(prRepo, userRepo, teamRepo)

	// иницилазиация хендлеров
	userHandler := handlers.NewUserHandler(userUseCase)
	teamHandler := handlers.NewTeamHandler(teamUseCase)
	prHandler := handlers.NewPRHandler(prUseCase)

	// инициализация роутера
	router := delivery.NewRouter(userHandler, teamHandler, prHandler)
	ginRouter := router.SetupRoutes()

	// запуск сервера
	log.Printf("запуск сервера на порту %s в %s", cfg.Port, cfg.Environment)
	if err := ginRouter.Run(":" + cfg.Port); err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
