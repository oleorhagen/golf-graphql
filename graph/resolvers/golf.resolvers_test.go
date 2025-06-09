package graph

import (
	"context"
	"testing"

	"github.com/oleorhagen/golf-graphql/graph/model"
)

func TestMutationResolver_CreatePlayer(t *testing.T) {
	resolver := &mutationResolver{&Resolver{}}

	player, err := resolver.CreatePlayer(context.Background(), model.NewPlayer{Name: "Test Player"})
	if err != nil {
		t.Errorf("CreatePlayer resolver returned an error: %v", err)
	}

	if player == nil {
		t.Error("CreatePlayer resolver returned a nil player")
	}

	if player.Name != "Test Player" {
		t.Errorf("CreatePlayer returned wrong name: got %s, want Test Player", player.Name)
	}
}

func TestPlayersOrderBy_Validation(t *testing.T) {
	// Test ordering enum values
	orderBy := model.PlayersOrderByNameAsc
	if !orderBy.IsValid() {
		t.Error("PlayersOrderByNameAsc should be valid")
	}

	// Test invalid ordering
	invalidOrder := model.PlayersOrderBy("INVALID")
	if invalidOrder.IsValid() {
		t.Error("Invalid order should not be valid")
	}

	// Test all valid values
	validOrders := []model.PlayersOrderBy{
		model.PlayersOrderByNameAsc,
		model.PlayersOrderByNameDesc,
		model.PlayersOrderByHandicapAsc,
		model.PlayersOrderByHandicapDesc,
		model.PlayersOrderByIDAsc,
		model.PlayersOrderByIDDesc,
	}

	for _, order := range validOrders {
		if !order.IsValid() {
			t.Errorf("Order %s should be valid", order)
		}
	}
}

func TestPlayerCondition_Structure(t *testing.T) {
	// Test condition structure
	condition := &model.PlayerCondition{
		Name:                stringPtr("test"),
		Handicap:            int32Ptr(10),
		HandicapGreaterThan: int32Ptr(5),
		HandicapLessThan:    int32Ptr(15),
	}

	if condition.Name == nil || *condition.Name != "test" {
		t.Error("PlayerCondition name not set correctly")
	}

	if condition.Handicap == nil || *condition.Handicap != 10 {
		t.Error("PlayerCondition handicap not set correctly")
	}

	if condition.HandicapGreaterThan == nil || *condition.HandicapGreaterThan != 5 {
		t.Error("PlayerCondition handicapGreaterThan not set correctly")
	}

	if condition.HandicapLessThan == nil || *condition.HandicapLessThan != 15 {
		t.Error("PlayerCondition handicapLessThan not set correctly")
	}
}

func stringPtr(s string) *string {
	return &s
}

func int32Ptr(i int32) *int32 {
	return &i
}
