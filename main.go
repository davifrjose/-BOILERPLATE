package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davifrjose/BOILERPLATE/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
		DB *database.Queries
	}

func main() {
	 feed, err := urlToFeed("https://wagslane.dev/index.xml")
	 if err != nil {
		log.Fatal(err)
	 }
	 fmt.Println(feed)
	
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable must be set")
	}

	dbURL := os.Getenv("CONNECTION_STRING")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	
	apiCfg := apiConfig{
		DB: dbQueries,
	}

	const collectionConcurrency = 10
	const collectionInterval = time.Minute
	go startScraping(dbQueries, collectionConcurrency, collectionInterval)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Post("/users", apiCfg.handlerUsersCreate)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerUserGetByAPIKey))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedsCreate))
	v1Router.Get("/feeds", apiCfg.handlerFeedsGet)
	v1Router.Post("/feed/follows", apiCfg.middlewareAuth(apiCfg.handlerFeedsFolowCreate))
	v1Router.Delete("/feed/follows/{id}", apiCfg.middlewareAuth(apiCfg.handlerFeedsFolowDelete))
	v1Router.Get("/feed/follows", apiCfg.middlewareAuth(apiCfg.handlerFeedsFollowFromUser))
	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerPostsGet))
	
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)


	r.Mount("/v1",v1Router)
	srv := &http.Server{
		Addr: ":" + port,
		Handler: r,
	}


	log.Printf("Serving on port: %s\n", port)
	
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	
}
