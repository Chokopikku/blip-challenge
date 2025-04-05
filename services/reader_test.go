package services

import (
	"os"
	"path/filepath"
	"testing"
)

func createTempCSV(t *testing.T, content string) string {
	t.Helper()
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create temp CSV file: %v", err)
	}

	return tmpFile
}

func TestReadCommits_ValidCSV(t *testing.T) {
	csvContent := `timestamp,user,repository,files,additions,deletions
1743862838,fabio,repo1,2,10,5
1743862938,rafael,repo2,3,15,10
`

	file := createTempCSV(t, csvContent)
	reader := NewCommitReader(file)

	commits, err := reader.ReadCommits()
	if err != nil {
		t.Fatalf("unexpected error reading CSV: %v", err)
	}

	if len(commits) != 2 {
		t.Fatalf("expected 2 commits, got %d", len(commits))
	}

	if commits[0].User != "fabio" || commits[1].Repository != "repo2" {
		t.Errorf("parsed data mismatch: %+v", commits)
	}
}

func TestReadCommits_FileNotFound(t *testing.T) {
	reader := NewCommitReader("nonexistent.csv")
	_, err := reader.ReadCommits()
	if err == nil {
		t.Error("expected error for nonexistent file, got nil")
	}
}

func TestReadCommits_InvalidNumbers(t *testing.T) {
	csvContent := `timestamp,user,repository,files,additions,deletions
notimestamp,fabio,repo1,x,y,?
`

	file := createTempCSV(t, csvContent)
	reader := NewCommitReader(file)

	commits, err := reader.ReadCommits()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(commits) != 1 {
		t.Fatalf("expected 1 commit, got %d", len(commits))
	}

	if commits[0].Timestamp != 0 || commits[0].Files != 0 {
		t.Errorf("expected defaulted zero values on parse error, got %+v", commits[0])
	}
}

func TestReadCommits_OnlyHeader(t *testing.T) {
	csvContent := `timestamp,user,repository,files,additions,deletions
`

	file := createTempCSV(t, csvContent)
	reader := NewCommitReader(file)

	commits, err := reader.ReadCommits()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(commits) != 0 {
		t.Errorf("expected 0 commits, got %d", len(commits))
	}
}
