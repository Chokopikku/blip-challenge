package services

import (
	"chokopikku/blip-challenge/models"
	"testing"
)

func TestActivityScorer_Calculate_Basic(t *testing.T) {
	scorer := NewActivityScorer(1.0, 0.4, 0.2)

	commit := models.Commit{
		Files:     5,
		Additions: 13,
		Deletions: 7,
	}

	expected := 1.0 + 0.4*5 + 0.2*20 // 1 + 2 + 4 = 7
	actual := scorer.Calculate(commit)

	if actual != expected {
		t.Errorf("expected %.2f, got %.2f", expected, actual)
	}
}

func TestActivityScorer_Calculate_ZeroWeights(t *testing.T) {
	scorer := NewActivityScorer(0, 0, 0)

	commit := models.Commit{
		Files:     5,
		Additions: 20,
		Deletions: 10,
	}

	if score := scorer.Calculate(commit); score != 0 {
		t.Errorf("expected score 0 with zero weights, got %.2f", score)
	}
}

func TestActivityScorer_Calculate_EmptyCommit(t *testing.T) {
	scorer := NewActivityScorer(2.0, 0.5, 0.1)

	commit := models.Commit{
		Files:     0,
		Additions: 0,
		Deletions: 0,
	}

	expected := 2.0
	actual := scorer.Calculate(commit)

	if actual != expected {
		t.Errorf("expected %.2f, got %.2f", expected, actual)
	}
}

func TestActivityScorer_Calculate_NegativeValues(t *testing.T) {
	scorer := NewActivityScorer(1.0, 1.0, 0.1)

	commit := models.Commit{
		Files:     -2,
		Additions: -5,
		Deletions: -3,
	}

	expected := -1.8 // 1 + (-2) + 0.1*(-8) = 1 - 2 - 0.8 = -1.8
	actual := scorer.Calculate(commit)

	if actual != expected {
		t.Errorf("expected %.2f, got %.2f", expected, actual)
	}
}
