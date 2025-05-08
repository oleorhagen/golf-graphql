package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oleorhagen/golf-graphql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	players     []*model.Player
	scorecards  []*model.Scorecard
	courses     []*model.Course
	tournaments []*model.Tournament
	DB          *pgxpool.Pool
}
