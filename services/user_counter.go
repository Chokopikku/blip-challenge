package services

import "chokopikku/blip-challenge/models"

// UserCounter keeps track of unique users for each repo.
type UserCounter struct {
	repoUsers map[string]map[string]struct{}
}

func NewUserCounter() *UserCounter {
	return &UserCounter{repoUsers: make(map[string]map[string]struct{})}
}

func (t *UserCounter) Add(commit models.Commit) {
	if commit.User == "" {
		return
	}
	if _, exists := t.repoUsers[commit.Repository]; !exists {
		t.repoUsers[commit.Repository] = make(map[string]struct{})
	}
	t.repoUsers[commit.Repository][commit.User] = struct{}{}
}

func (t *UserCounter) GetUniqueUserCount(repo string) int {
	return len(t.repoUsers[repo])
}
