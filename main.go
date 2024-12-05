package main

import (
	"database/sql"
	"github.com/ArtiomStartev/rss-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	ServerPort       = "PORT"
	DatabaseURL      = "DB_URL"
	PostgreSqlDriver = "postgres"
	EnvFileName      = ".env"
	ConcurrentFeeds  = 10
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	if err := godotenv.Load(EnvFileName); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv(ServerPort)
	if port == "" {
		log.Fatal("Server port is not found in the environment")
	}

	dbURL := os.Getenv(DatabaseURL)
	if dbURL == "" {
		log.Fatal("Database URL is not found in the environment")
	}

	conn, err := sql.Open(PostgreSqlDriver, dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	go startScraping(apiCfg.DB, ConcurrentFeeds, time.Minute)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	// GET Requests
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUserByAPIKey))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1Router.Get("/feed-follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Get("/user-posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))

	// POST Requests
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Post("/feed-follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))

	// DELETE Requests
	v1Router.Delete("/feed-follows/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	if err = server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
