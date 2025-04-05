package utils

import (
	"chokopikku/blip-challenge/models"
	"encoding/json"
	"os"
)

func ExportRankingAsJSON(ranking []models.RepositoryScore, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(ranking)
}
