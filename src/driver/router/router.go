package router

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
	"github.com/yuorei/bukubukubooking-back/graph/generated"
	"github.com/yuorei/bukubukubooking-back/src/adapter/infrastructure"
	"github.com/yuorei/bukubukubooking-back/src/adapter/presentation/resolver"
	"github.com/yuorei/bukubukubooking-back/src/application"
	"github.com/yuorei/bukubukubooking-back/src/middleware"
)

func NewRouter() {
	const defaultPort = "8080"

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	infra := infrastructure.NewInfrastructure()
	app := application.NewApplication(infra)
	r := resolver.NewResolver(app)
	c := generated.Config{Resolvers: r}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		Logger:         log.New(os.Stdout, "bukubukubooking-server", log.LstdFlags),
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders: []string{"*"},
	})

	middleware.Middleware()
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", corsOpts.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
