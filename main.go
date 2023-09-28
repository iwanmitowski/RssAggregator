package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/iwanmitowski/RssAggregator/internal/database"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB Database
}

type Database interface {
	// Define the methods you need for your application
}

type PostGresDBClient struct {
	*database.Queries
}

type MongoDBClient struct {
	*mongo.Client
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT not found in .env")
	}

	// To apply migrations and build the queries from sql/queries
	// open CMD in the project on sqlc.yaml level and run the following code:
	// docker pull sqlc/sqlc
	// docker run --rm -v "%cd%:/src" -w /src sqlc/sqlc generate
	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Printf("DB_URL for Postgres not found in .env")
	}

	dbUrl = os.Getenv("DB_URL_MONGO")
	isMongo := false

	if dbUrl == "" {
		log.Fatal("DB_URL not found in .env")
	} else {
		isMongo = true
	}

	var db interface{}
	if isMongo {
		mongoClient := connectToMongo(dbUrl)
		defer func() {
			if mongoClient != nil {
				mongoClient.Disconnect(nil)
			}
		}()
		db = mongoClient
	} else {
		// Use your SQL database connection here
		conn, err := sql.Open("postgres", dbUrl)
		if err != nil {
			log.Fatal("Can't connect to db.")
		}
		db = &PostGresDBClient{Queries: database.New(conn)}
	}

	apiCfg := apiConfig{
		DB: db,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Post("/register", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))
	v1Router.Get("/feeds", apiCfg.middlewareAuth(apiCfg.handlerGetNotFollowedFeeds))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Post("/feed/follow", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed/followed", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed/unfollow/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerUnfollowFeed))

	router.Mount("/v1", v1Router)

	log.Printf("Server starting on port %v", port)

	// New routine so it doesn't affect the flow

	go startScraping(apiCfg.DB, 10, time.Minute)

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func connectToMongo(dbURLMongo string) *MongoDBClient {
	clientOptions := options.Client().ApplyURI(dbURLMongo)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	return &MongoDBClient{Client: client}
}
