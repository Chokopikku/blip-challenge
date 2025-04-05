package services

import (
	"chokopikku/blip-challenge/models"
	"math"
	"sort"
)

// RepositoryRanker finds the top 10 most active repos based on their score.
type RepositoryRanker struct{}

func NewRepositoryRanker() *RepositoryRanker {
	return &RepositoryRanker{}
}

func (r *RepositoryRanker) Rank(commits []models.Commit, scorer *ActivityScorer, userTracker *UserCounter) []models.RepositoryScore {
	repoScores := make(map[string]float64)

	for _, commit := range commits {
		score := scorer.Calculate(commit)
		repoScores[commit.Repository] += score
	}

	var ranking []models.RepositoryScore
	for repo, score := range repoScores {
		uniqueUsers := userTracker.GetUniqueUserCount(repo)
		finalScore := score * (1 + math.Log(1+float64(uniqueUsers)))

		ranking = append(ranking, models.RepositoryScore{Name: repo, Score: finalScore})
	}

	sort.Slice(ranking, func(i, j int) bool {
		return ranking[i].Score > ranking[j].Score
	})

	if len(ranking) > 10 {
		ranking = ranking[:10]
	}

	return ranking
}
