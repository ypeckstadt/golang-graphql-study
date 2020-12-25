package graph

import (
	"github.com/ypeckstadt/golang-graphql-study/pkg/server/user"
)

type Resolver struct {
	User user.Service
}
