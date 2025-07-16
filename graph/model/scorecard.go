package model

import (
	"time"
	"github.com/google/uuid"
)

type Scorecard struct {
	ID           uuid.UUID  `json:"id"`
	TournamentID *uuid.UUID `json:"tournament_id,omitempty"`
	Handicap     int32      `json:"handicap"`
	CreatedAt    time.Time  `json:"created_at"`
	CourseName   string     `json:"course_name"`
	PlayerID     uuid.UUID
	Player       *Player `json:"player"`
}
