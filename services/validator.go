package services

import (
	"chokopikku/blip-challenge/models"
	"errors"
)

// CommitValidator validates commit data.
type CommitValidator struct{}

func NewCommitValidator() *CommitValidator {
	return &CommitValidator{}
}

func (v *CommitValidator) Validate(commit models.Commit, usernameRequired bool) error {
	if commit.Repository == "" {
		return errors.New("repository name is empty")
	}
	if usernameRequired && commit.User == "" {
		return errors.New("user name is empty")
	}
	if commit.Files < 0 || commit.Additions < 0 || commit.Deletions < 0 {
		return errors.New("invalid data: negative values")
	}
	return nil
}
