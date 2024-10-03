package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/appointment-api/graph"
	"github.com/appointment-api/middleware"
	"github.com/appointment-api/service"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	services := service.NewServiceContainer()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			Services: services,
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.DummyMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
