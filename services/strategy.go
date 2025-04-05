package services

import (
	"math"
)

type ScoringStrategy interface {
	Score(repo string, baseScore float64, userCounter *UserCounter) float64
}

// BasicStrategy uses the formula: 1*commits + x*files_changed + y*(line_addition + line_deletions)
type BasicStrategy struct{}

func (s *BasicStrategy) Score(repo string, baseScore float64, _ *UserCounter) float64 {
	return baseScore
}

// UserWeightedStrategy uses the formula: base_score*(1 + ln(1 + unique_users))
type UserWeightedStrategy struct{}

func (s *UserWeightedStrategy) Score(repo string, baseScore float64, userCounter *UserCounter) float64 {
	users := userCounter.GetUniqueUserCount(repo)
	return baseScore * (1 + math.Log(1+float64(users)))
}
