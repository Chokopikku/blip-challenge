package models

// Commit represents a single commit record from the CSV file.
type Commit struct {
	Timestamp  int64
	User       string
	Repository string
	Files      int
	Additions  int
	Deletions  int
}

// RepositoryScore represents a repository and its calculated activity score.
type RepositoryScore struct {
	Name  string
	Score float64
}

/*
Fields are aligned based on their size.
The compiler will not add memory padding between them.
*/
