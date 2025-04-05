package services

import (
	"chokopikku/blip-challenge/models"
	"testing"
)

func TestUserCounter_AddAndCount(t *testing.T) {
	counter := NewUserCounter()

	commits := []models.Commit{
		{User: "fabio", Repository: "repo1"},
		{User: "rafael", Repository: "repo1"},
		{User: "fabio", Repository: "repo1"},
		{User: "pereira", Repository: "repo2"},
		{User: "", Repository: "repo2"},
	}

	for _, c := range commits {
		counter.Add(c)
	}

	if count := counter.GetUniqueUserCount("repo1"); count != 2 {
		t.Errorf("expected 2 unique users for repo1, got %d", count)
	}

	if count := counter.GetUniqueUserCount("repo2"); count != 1 {
		t.Errorf("expected 1 unique user for repo2, got %d", count)
	}
}

func TestUserCounter_UnknownRepo(t *testing.T) {
	counter := NewUserCounter()

	if count := counter.GetUniqueUserCount("nonexistent"); count != 0 {
		t.Errorf("expected 0 users for unknown repo, got %d", count)
	}
}
