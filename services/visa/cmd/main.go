package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/marelinaa/visa-api/services/visa/internal/config"
	"github.com/marelinaa/visa-api/services/visa/internal/handler/apply"
	"github.com/marelinaa/visa-api/services/visa/internal/handler/auth"
	"github.com/marelinaa/visa-api/services/visa/internal/redisconn"
	"github.com/marelinaa/visa-api/services/visa/internal/repository/psql"
	cache "github.com/marelinaa/visa-api/services/visa/internal/repository/redis"
	applySrv "github.com/marelinaa/visa-api/services/visa/internal/service/apply"
	authSrv "github.com/marelinaa/visa-api/services/visa/internal/service/auth"

	"github.com/marelinaa/visa-api/services/visa/migrations"
)

type Repos struct {
	psql  psql.UserRepo
	redis cache.Repo
}

type Services struct {
	auth  *authSrv.AuthService
	apply *applySrv.ApplicantService
}

type Handlers struct {
	auth  auth.AuthHandler
	apply apply.ApplyHandler
}

func CreateHandlers(srv *Services) *Handlers {
	handlers := &Handlers{
		auth: auth.NewAuthHandler(srv.auth),
	}

	return handlers
}

func CreateServices(psqlRepo *psql.UserRepo, redisRepo *cache.Repo) *Services {
	authSrv := authSrv.NewAuthService(psqlRepo, redisRepo)
	//applySrv := applySrv.NewApplicantService(psqlRepo)

	srv := &Services{
		auth: authSrv,
	}

	return srv
}

func CreateRepos(redisClient *redis.Client, db *sql.DB) *Repos {
	redisRepo := cache.NewRepository(redisClient)
	psqlRepo := psql.NewRepository(db)

	repos := &Repos{
		psql:  *psqlRepo,
		redis: *redisRepo,
	}

	return repos
}

func ConnectRedis(config *config.Config) (*redis.Client, error) {
	redisAddr := fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)
	redisClient := redisconn.NewRedisClient(redisAddr, config.Redis.Password, config.Redis.DB)
	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}

	return redisClient, err
}

func ConnectPostgres(config *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		config.DB.User, config.DB.Password,
		config.DB.DBName, config.DB.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	//run migrations
	err := migrations.RunMigrations()
	if err != nil {
		log.Fatal("failed to run migrations: ", err)
	}

	config, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// connecting to Redis and PostgreSQL
	redisClient, err := ConnectRedis(config)
	if err != nil {
		log.Fatalf("could not connect to Redis: %v", err)
	}

	db, err := ConnectPostgres(config)
	if err != nil {
		log.Fatalf("could not connect to Postgres: %v", err)
	}

	//create repository instances
	repos := CreateRepos(redisClient, db)

	//create service instances
	srv := CreateServices(&repos.psql, &repos.redis)

	// create handler instances
	handlers := CreateHandlers(srv)

	router := gin.Default()

	// Обслуживание статических файлов
	router.Static("/static", "./services/visa/static")

	// Главная страница
	router.GET("/v1/visa/apply", func(c *gin.Context) {
		c.File("./services/visa/static/index.html")
	})

	// Определение маршрутов
	handlers.apply.DefineRoutes(router)

	apiPort := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	log.Printf("Starting server on %s\n", apiPort)
	log.Fatal(http.ListenAndServe(apiPort, router))
}
