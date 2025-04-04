package services

import "chokopikku/blip-challenge/models"

// ActivityScorer calculates the activity score for a repository.
type ActivityScorer struct{}

func NewActivityScorer() *ActivityScorer {
	return &ActivityScorer{}
}

func (c *ActivityScorer) Calculate(commit models.Commit) float64 {
	return 1 + 0.5*float64(commit.Files) + 0.2*float64(commit.Additions+commit.Deletions)
}

// TODO add username count to formula
