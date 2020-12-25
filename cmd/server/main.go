package main

import (
	"github.com/ypeckstadt/go-env"
	"github.com/ypeckstadt/go-utils"
	"github.com/ypeckstadt/golang-graphql-study/pkg/server"
	"github.com/ypeckstadt/golang-graphql-study/pkg/utils/environment"
)

func main() {
	// Load ENV variables
	var environment environment.Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	utils.PanicWhenHasError(err)

	// Start API service
	utils.PanicWhenHasError(server.Start(environment))
}
