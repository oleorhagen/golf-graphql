package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/oleorhagen/golf-graphql/graph/model"

	"github.com/jackc/pgx/v5"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	players     []*model.Player
	tournaments []*model.Tournament
	// todos []*model.Todo
	// users []*model.User
	DB *pgx.Conn
}
