package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/lululululu5/blog-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DB URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Could not initiate db connection")
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	} 


	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/users", apiCfg.handlerUserCreate)
	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.handlerUserGetByAPI))

	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	mux.HandleFunc("GET /v1/feeds", apiCfg.handlerGetFeeds)
	
	
	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerError)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}



