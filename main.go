package main

import (
	"chokopikku/blip-challenge/services"
	"chokopikku/blip-challenge/utils"
	"fmt"
	"log"
)

func main() {
	utils.SetupLogger()

	reader := services.NewCommitReader("commits.csv")
	commits, err := reader.ReadCommits()
	if err != nil {
		log.Fatalf("Error reading commits: %v", err)
	}

	validator := services.NewCommitValidator()
	for _, commit := range commits {
		if err := validator.Validate(commit, false); err != nil {
			log.Printf("Invalid commit: %v", err)
			continue
		}
	}

	scorer := services.NewActivityScorer()
	ranker := services.NewRepositoryRanker()
	ranking := ranker.Rank(commits, scorer)

	fmt.Println("Top 10 Most Active Repositories:")
	for i, repo := range ranking {
		fmt.Printf("%d. %s - Activity Score: %.2f\n", i+1, repo.Name, repo.Score)
	}
}
