package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ypeckstadt/go-utils"
	"github.com/ypeckstadt/golang-graphql-study/pkg/graph/generated"
	gqlmodel "github.com/ypeckstadt/golang-graphql-study/pkg/graph/model"
	graph "github.com/ypeckstadt/golang-graphql-study/pkg/graph/resolver"
	"github.com/ypeckstadt/golang-graphql-study/pkg/server/user"
	"github.com/ypeckstadt/golang-graphql-study/pkg/utils/environment"
	"github.com/ypeckstadt/golang-graphql-study/pkg/utils/postgres"
	"github.com/ypeckstadt/golang-graphql-study/pkg/utils/uuid/satori"
	"log"
	"net/http"
)

// Start starts the client application
func Start(env environment.Environment) error {


	// Connect to postgres database
	db, err := postgres.New(postgres.ConnectionSettings{
		Host:     env.Database.Host,
		Port:     env.Database.Port,
		Username: env.Database.Username,
		Password: env.Database.Password,
		Name:     env.Database.Name,
	})
	utils.PanicWhenHasError(err)


	// create services
	satoriUUIDService := satori.New()
	userService := user.Initialize(db, satoriUUIDService)

	// setup Graph QL
	config := generated.Config{Resolvers: &graph.Resolver{User: &userService}}
	//config.Directives.HasRole = hasRoleDirective
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", env.Server.Port)
	log.Fatal(http.ListenAndServe(":"+env.Server.Port, nil))

	return nil
}

func hasRoleDirective(ctx context.Context, obj interface{}, next graphql.Resolver, role gqlmodel.Role) (res interface{}, err error) {
	log.Printf("Inside hasRoleDirective - ignore the role check for now")
	return next(ctx)
}
