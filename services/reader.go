package services

import (
	"encoding/csv"
	"os"
	"strconv"

	"chokopikku/blip-challenge/models"
)

// CommitReader reads commits from a CSV file.
type CommitReader struct {
	FilePath string
}

func NewCommitReader(filePath string) *CommitReader {
	return &CommitReader{FilePath: filePath}
}

func (r *CommitReader) ReadCommits() ([]models.Commit, error) {
	file, err := os.Open(r.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var commits []models.Commit
	for _, record := range records[1:] { // Skip header
		timestamp, _ := strconv.ParseInt(record[0], 10, 64)
		files, _ := strconv.Atoi(record[3])
		additions, _ := strconv.Atoi(record[4])
		deletions, _ := strconv.Atoi(record[5])

		commit := models.Commit{
			Timestamp:  timestamp,
			User:       record[1],
			Repository: record[2],
			Files:      files,
			Additions:  additions,
			Deletions:  deletions,
		}

		commits = append(commits, commit)
	}

	return commits, nil
}
