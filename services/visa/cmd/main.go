package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marelinaa/visa-api/services/visa/internal/handler"
	"github.com/marelinaa/visa-api/services/visa/internal/repository"
	"github.com/marelinaa/visa-api/services/visa/internal/service"
	"github.com/marelinaa/visa-api/services/visa/migrations"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dsn := os.Getenv("DB_URL")
	err = migrations.RunMigrations(dsn)
	if err != nil {
		log.Fatal("failed to run migrations: ", err)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	applicantService := service.NewApplicantService(repo)
	applicantHandler := handler.NewApplicantHandler(applicantService)

	router := gin.Default()

	// Обслуживание статических файлов
	router.Static("/static", "./services/visa/static")

	// Главная страница
	router.GET("/v1/visa/apply", func(c *gin.Context) {
		c.File("./services/visa/static/index.html")
	})

	// Определение маршрутов
	applicantHandler.DefineRoutes(router)

	apiPort := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	log.Printf("Starting server on %s\n", apiPort)
	log.Fatal(http.ListenAndServe(apiPort, router))
}
