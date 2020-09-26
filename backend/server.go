package main

import (
	"gql/graph"
	"gql/graph/generated"
	"log"
	"gql/postgres"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v10"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "5000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	pgDB := postgres.New(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "openclosed",
		Database: "YourJube",
	})

	// pgDB := pg.Connect(&pg.Options{
		// Addr:     ":5432",
		// User:     "postgres",
		// Password: "openclosed",
		// Database: "YourJube",
	// })

	
	pgDB.AddQueryHook(postgres.DBLogger{})

	defer pgDB.Close()

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: pgDB}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	// log.Println(db)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
