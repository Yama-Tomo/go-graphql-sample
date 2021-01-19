package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"sample/graph/generated"
	"sample/graph/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-sql-driver/mysql"
	"sample/ent"
	"sample/ent/migrate"
)

const defaultPort = "8080"

func main() {
	loadEnv()
	entOptions := []ent.Option{}

	if os.Getenv("GO_ENV") != "production" {
		entOptions = append(entOptions, ent.Debug())
	}

	mc := mysql.Config{
		User:                 os.Getenv("DATABASE_USER"),
		Passwd:               os.Getenv("DATABASE_PASS"),
		Net:                  os.Getenv("DATABASE_NET"),
		Addr:                 os.Getenv("DATABASE_ADDR"),
		DBName:               os.Getenv("DATABASE_NAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	client, err := ent.Open("mysql", mc.FormatDSN(), entOptions...)
	if err != nil {
		log.Fatalf("Error open mysql ent client: %v\n", err)
	}

	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithDropIndex(true),
		migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{DB: client}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed load `.env`: [%s]", err)
	}
}
