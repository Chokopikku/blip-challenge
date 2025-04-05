package services

import "chokopikku/blip-challenge/models"

// ActivityScorer calculates the activity score for a repository.
type ActivityScorer struct {
	weightCommits float64
	weightFiles   float64
	weightLines   float64
}

func NewActivityScorer(commits, files, lines float64) *ActivityScorer {
	return &ActivityScorer{
		weightCommits: commits,
		weightFiles:   files,
		weightLines:   lines,
	}
}

func (s *ActivityScorer) Calculate(commit models.Commit) float64 {
	return s.weightCommits +
		s.weightFiles*float64(commit.Files) +
		s.weightLines*float64(commit.Additions+commit.Deletions)
}
