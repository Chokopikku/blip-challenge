package services

import (
	"chokopikku/blip-challenge/models"
	"math"
	"testing"
)

func TestBasicStrategy_Score(t *testing.T) {
	strategy := &BasicStrategy{}
	userCounter := NewUserCounter()

	baseScore := 42.0
	finalScore := strategy.Score("repo1", baseScore, userCounter)

	if finalScore != baseScore {
		t.Errorf("expected %.2f, got %.2f", baseScore, finalScore)
	}
}

func TestUserWeightedStrategy_WithUsers(t *testing.T) {
	userCounter := NewUserCounter()
	userCounter.Add(models.Commit{Repository: "repo1", User: "alice"})
	userCounter.Add(models.Commit{Repository: "repo1", User: "bob"})
	userCounter.Add(models.Commit{Repository: "repo1", User: "charlie"})

	baseScore := 10.0
	expectedMultiplier := 1 + math.Log(1+3)
	expected := baseScore * expectedMultiplier

	strategy := &UserWeightedStrategy{}
	actual := strategy.Score("repo1", baseScore, userCounter)

	if math.Abs(expected-actual) > 0.0001 {
		t.Errorf("expected %.4f, got %.4f", expected, actual)
	}
}

func TestUserWeightedStrategy_NoUsers(t *testing.T) {
	userCounter := NewUserCounter()

	baseScore := 10.0
	expected := baseScore * (1 + math.Log(1)) // 10 * (1 + ln(1)) = 10 * (1 + 0) = 10

	strategy := &UserWeightedStrategy{}
	actual := strategy.Score("repo1", baseScore, userCounter)

	if actual != expected {
		t.Errorf("expected %.2f, got %.2f", expected, actual)
	}
}
