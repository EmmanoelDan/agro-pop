package routers

import (
	"fmt"
	"log"
	"os"

	"github.com/EmmanoelDan/agro-pop/handlers"
	"github.com/EmmanoelDan/agro-pop/repositories"
	"github.com/EmmanoelDan/agro-pop/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	r := gin.Default()
	initializeRouter(r)

	port := ":8080"
	log.Printf("Server is listening on port %v", port)

	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func initializeRouter(r *gin.Engine) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
        log.Fatal("Error connecting to the database", err)
    }

	userRepo := &repositories.UserRepository{DB: db}

	registerUserUsecase := usecases.NewRegisterUserUseCase(userRepo)
	registerUserHandler := handlers.NewRegisterUserHandler(registerUserUsecase)

	authUsecase := usecases.NewAuthUseCase(*userRepo)
	authHandler := handlers.NewAuthHandler(authUsecase)

	r.POST("/register", registerUserHandler.Register)
	r.POST("/login", authHandler.Login)
}