package graph

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	pgx "github.com/jackc/pgx/v5"
	"github.com/oleorhagen/golf-graphql/graph/model"
)

func getScorecards(r *playerResolver, ctx context.Context, scorerID uuid.UUID) ([]*model.Scorecard, error) {
	var scorecards []*model.Scorecard
	var id uuid.UUID
	var tournamentID uuid.UUID
	var playerID uuid.UUID
	var handicap int32
	var course_name string
	fmt.Fprintf(os.Stderr, "Querying scorecard with playerID=%s\n", scorerID)
	rows, err := r.DB.Query(ctx,
		"select id, tournament_id, scorer_id, handicap, course_name from scorecard where scorer_id=$1",
		scorerID,
	)
	_, err = pgx.ForEachRow(rows, []any{&id, &tournamentID, &playerID, &handicap, &course_name}, func() error {
		fmt.Fprintf(os.Stderr, "Got: %v\n", id)
		scorecards = append(scorecards, &model.Scorecard{
			ID:           id,
			TournamentID: tournamentID,
			Handicap:     handicap,
			CourseName:   course_name,
			PlayerID:     playerID,
		})
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, fmt.Errorf("Failed to query the database for scorecards: %w", err)
	}

	r.scorecards = append(r.scorecards, scorecards...)
	return r.scorecards, nil
}
