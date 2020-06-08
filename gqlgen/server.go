package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MisakiFx/Golang/gqlgen/graph/generated"
	"github.com/MisakiFx/Golang/gqlgen/graph/resolvers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//这块是原生的http接口
	//srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)
	//下面使用gin替代
	server := gin.New()
	server.Any("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	server.Any("/", func(context *gin.Context) {
		request := context.Request
		response := context.Writer
		playground.Handler("GraphQL playground", "/query")(response, request)
	})
	server.Any("/query", func(context *gin.Context) {
		request := context.Request
		response := context.Writer
		handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}})).ServeHTTP(response, request)
	})
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(server.Run(":8080"))
}
