package main

import (
	"log"
	"os"
	"strconv"

	"github.com/arsmn/fastgql/graphql/handler"
	"github.com/arsmn/fastgql/graphql/playground"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gopato.io/db"
	"gopato.io/graph"
	"gopato.io/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("cant load .env file")
	}
	prod, _ := strconv.ParseBool(os.Getenv("PROD"))
	app := fiber.New(
		fiber.Config{
			Prefork: prod,
			AppName: "gopatogql",
		},
	)
	db.ConnectDb()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	gqlHandler := srv.Handler()
	playground := playground.Handler("GraphQL playground", "/query")

	app.All("/query", func(c *fiber.Ctx) error {
		gqlHandler(c.Context())
		return nil
	})

	app.All("/", func(c *fiber.Ctx) error {
		playground(c.Context())
		return nil
	})
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(app.Listen(":" + port))
}
