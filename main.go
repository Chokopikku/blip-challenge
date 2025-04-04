package main

import (
	"chokopikku/blip-challenge/services"
	"chokopikku/blip-challenge/utils"
	"fmt"
)

func main() {
	logger, err := utils.NewLogger("app.log")
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		return
	}
	defer logger.Close()

	logger.Info("Application started")

	reader := services.NewCommitReader("commits.csv")
	commits, err := reader.ReadCommits()
	if err != nil {
		logger.Error(fmt.Sprintf("Error reading commits: %v", err))
		return
	}

	logger.Info("Data reading completed")

	validator := services.NewCommitValidator()
	for _, commit := range commits {
		if err := validator.Validate(commit, false); err != nil {
			logger.Warn(fmt.Sprintf("Invalid commit: %v", err))
			continue
		}
	}

	logger.Info("Data validation completed")

	scorer := services.NewActivityScorer()
	ranker := services.NewRepositoryRanker()
	ranking := ranker.Rank(commits, scorer)

	logger.Info("Ranking calculation completed")

	fmt.Println("Top 10 Most Active Repositories:")
	for i, repo := range ranking {
		fmt.Printf("%d. %s - Activity Score: %.2f\n", i+1, repo.Name, repo.Score)
		logger.Debug(fmt.Sprintf("%d. %s - Activity Score: %.2f\n", i+1, repo.Name, repo.Score))
	}

	logger.Info("Application finished")
}
