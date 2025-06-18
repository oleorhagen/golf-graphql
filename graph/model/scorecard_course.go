package model

import "github.com/google/uuid"

type ScorecardCourse struct {
	Name         string           `json:"name"`
	Slope        int32            `json:"slope"`
	CourseRating float64          `json:"course_rating"`
	NrHoles      int32            `json:"nr_holes"`
	Holes        []*ScorecardHole `json:"holes"`
	ScorecardID  uuid.UUID        // Not exposed in GraphQL, used for resolving holes
}