package services

import (
	"chokopikku/blip-challenge/models"
	"testing"
)

func TestCommitValidator_ValidCommit(t *testing.T) {
	validator := NewCommitValidator()
	commit := models.Commit{
		User:       "alice",
		Repository: "repo1",
		Files:      3,
		Additions:  10,
		Deletions:  2,
	}
	err := validator.Validate(commit, true)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestCommitValidator_MissingRepository(t *testing.T) {
	validator := NewCommitValidator()
	commit := models.Commit{
		User:      "alice",
		Files:     2,
		Additions: 5,
	}
	err := validator.Validate(commit, true)
	if err == nil || err.Error() != "repository name is empty" {
		t.Errorf("expected repository name error, got %v", err)
	}
}

func TestCommitValidator_MissingUser_Required(t *testing.T) {
	validator := NewCommitValidator()
	commit := models.Commit{
		Repository: "repo1",
		Files:      1,
		Additions:  5,
		Deletions:  3,
	}
	err := validator.Validate(commit, true)
	if err == nil || err.Error() != "user name is empty" {
		t.Errorf("expected username error, got %v", err)
	}
}

func TestCommitValidator_MissingUser_NotRequired(t *testing.T) {
	validator := NewCommitValidator()
	commit := models.Commit{
		Repository: "repo1",
		Files:      1,
		Additions:  5,
		Deletions:  3,
	}
	err := validator.Validate(commit, false)
	if err != nil {
		t.Errorf("expected no error for optional username, got %v", err)
	}
}

func TestCommitValidator_NegativeValues(t *testing.T) {
	validator := NewCommitValidator()
	commit := models.Commit{
		User:       "bob",
		Repository: "repo2",
		Files:      -1,
		Additions:  5,
		Deletions:  3,
	}
	err := validator.Validate(commit, true)
	if err == nil || err.Error() != "invalid data: negative values" {
		t.Errorf("expected negative value error, got %v", err)
	}
}
