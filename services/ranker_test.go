package services

import (
	"chokopikku/blip-challenge/models"
	"strconv"
	"testing"
)

func TestRank_WithBasicStrategy(t *testing.T) {
	commits := []models.Commit{
		{Repository: "repo1", User: "fabio", Files: 1, Additions: 10, Deletions: 2},
		{Repository: "repo1", User: "rafael", Files: 2, Additions: 5, Deletions: 1},
		{Repository: "repo1", User: "rafael", Files: 3, Additions: 20, Deletions: 5},
		{Repository: "repo2", User: "teixeira", Files: 3, Additions: 20, Deletions: 5},
		{Repository: "repo2", User: "pereira", Files: 3, Additions: 20, Deletions: 5},
		{Repository: "repo3", User: "fabio", Files: 3, Additions: 20, Deletions: 5},
	}

	userCounter := NewUserCounter()
	for _, commit := range commits {
		userCounter.Add(commit)
	}

	scorer := NewActivityScorer(1.0, 0.5, 0.2)
	strategy := &BasicStrategy{}

	ranker := NewRepositoryRanker()
	result := ranker.Rank(commits, scorer, userCounter, strategy)

	if len(result) != 3 {
		t.Fatalf("expected 3 repos in ranking, got %d", len(result))
	}

	if result[0].Name != "repo2" {
		t.Errorf("expected repo2 to be ranked first, got %s", result[0].Name)
	}
}

func TestRank_WithUserWeightedStrategy(t *testing.T) {
	commits := []models.Commit{
		{Repository: "repo1", User: "fabio", Files: 1, Additions: 10, Deletions: 2},
		{Repository: "repo1", User: "rafael", Files: 2, Additions: 5, Deletions: 1},
		{Repository: "repo1", User: "rafael", Files: 3, Additions: 20, Deletions: 5},
		{Repository: "repo2", User: "teixeira", Files: 3, Additions: 20, Deletions: 5},
		{Repository: "repo2", User: "teixeira", Files: 3, Additions: 20, Deletions: 5},
		{Repository: "repo3", User: "pereira", Files: 3, Additions: 20, Deletions: 5},
	}

	userCounter := NewUserCounter()
	for _, commit := range commits {
		userCounter.Add(commit)
	}

	scorer := NewActivityScorer(1.0, 0.5, 0.2)
	strategy := &UserWeightedStrategy{}

	ranker := NewRepositoryRanker()
	result := ranker.Rank(commits, scorer, userCounter, strategy)

	if len(result) != 3 {
		t.Fatalf("expected 3 repos in ranking, got %d", len(result))
	}

	if result[0].Name != "repo1" {
		t.Errorf("expected repo1 to be ranked first due to user weighting, got %s", result[0].Name)
	}
}

func TestRank_NoCommits(t *testing.T) {
	var commits []models.Commit

	userCounter := NewUserCounter()
	scorer := NewActivityScorer(1.0, 0.5, 0.2)
	strategy := &BasicStrategy{}
	ranker := NewRepositoryRanker()

	result := ranker.Rank(commits, scorer, userCounter, strategy)

	if len(result) != 0 {
		t.Errorf("expected empty ranking, got %d entries", len(result))
	}
}

func TestRank_LimitToTop10(t *testing.T) {
	var commits []models.Commit

	for i := 0; i < 15; i++ {
		repoName := "repo" + strconv.Itoa(i)
		commits = append(commits, models.Commit{
			Repository: repoName,
			User:       "user1",
			Files:      1,
			Additions:  1,
			Deletions:  0,
		})
	}

	userCounter := NewUserCounter()
	for _, c := range commits {
		userCounter.Add(c)
	}

	scorer := NewActivityScorer(1.0, 0.5, 0.2)
	strategy := &BasicStrategy{}
	ranker := NewRepositoryRanker()

	result := ranker.Rank(commits, scorer, userCounter, strategy)

	if len(result) != 10 {
		t.Errorf("expected top 10 results, got %d", len(result))
	}
}

func TestRank_BlankUsernamesAreIgnored(t *testing.T) {
	commits := []models.Commit{
		{Repository: "repoX", User: "", Files: 3, Additions: 10, Deletions: 0},
		{Repository: "repoX", User: "bob", Files: 1, Additions: 5, Deletions: 2},
		{Repository: "repoX", User: "", Files: 2, Additions: 6, Deletions: 1},
	}

	userCounter := NewUserCounter()
	for _, c := range commits {
		userCounter.Add(c)
	}

	count := userCounter.GetUniqueUserCount("repoX")
	if count != 1 {
		t.Errorf("expected 1 unique user, got %d", count)
	}
}
