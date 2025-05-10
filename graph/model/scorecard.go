package model

import "github.com/google/uuid"

type Scorecard struct {
	ID           uuid.UUID `json:"id"`
	TournamentID uuid.UUID `json:"tournament_id"`
	Handicap     int32     `json:"handicap"`
	CourseName   string    `json:"course_name"`
	PlayerID     uuid.UUID
	Player       *Player `json:"player"`
}
