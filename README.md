# Repository Activity Scoring Algorithm
## Overview
This algorithm calculates an activity score for repositories based on commit data from CSV files. The algorithm allows customization through a variety of configuration options, including scoring strategies that account for the number of commits, files changed, and line additions/deletions. Additionally, the algorithm can factor in the number of unique users contributing to each repository.

## Key Features
- **Commit Scoring**: Scores repositories based on the number of commits, files changed, and lines added/deleted.
- **User-Weighted Scoring**: Optionally adjusts the score based on the number of unique contributors to a repository.
- **Flexible Strategy**: Supports multiple scoring strategies, which can be switched dynamically.
- **Top 10 Ranking**: Ranks repositories by their activity score, showing the top 10 most active repositories.

## Algorithm Breakdown
### 1. Commit Scoring
The activity score for each commit is calculated using the formula:

`score = weight_commits + (weight_files * files_changed) + (weight_lines * (lines_added + lines_deleted))`

Where:
- `weight_commits`, `weight_files`, and `weight_lines` are user-defined constants.
- `files_changed` is the number of files modified in the commit.
- `lines_added` and `lines_deleted` are the respective number of added and deleted lines in the commit.

### 2. Scoring Strategies
Two scoring strategies are provided:
- **Basic Strategy**: Simply sums up the commit score using the above formula without any adjustments.
- **User-Weighted Strategy**: Adjusts the score by factoring in the number of unique users for a repository. The formula is:
```
score = base_score * (1 + ln(1 + unique_users))
```
Where `unique_users` is the count of unique contributors to the repository.

### 3. Top 10 Ranking
Repositories are ranked based on their final activity score. The top 10 most active repositories are shown in the results.

## Setup and Running the Algorithm
### 1. Prerequisites
Ensure you have the following installed:
- Go (version 1.18+)
- Git (for managing dependencies)
- A CSV file containing commit data from source controlled repositories. The file must have the following columns:
    - `timestamp`: Unix timestamp of the commit
    - `user`: GitHub username of the commit author (can be empty)
    - `repository`: The name of the repository
    - `files`: The number of files changed in the commit
    - `additions`: The number of lines added in the commit
    - `deletions`: The number of lines deleted in the commit

### 2. Environment Configuration
Create an `.env` file in the project root with the following configuration variables:
```
LOG_FILE_PATH=logs/application.log       # Log file path
WEIGHT_COMMITS=1.0                       # Weight for commit count
WEIGHT_FILES=0.4                         # Weight for file changes
WEIGHT_LINES=0.2                         # Weight for line additions/deletions
```
You can adjust the weights for commit count, file changes, and line additions/deletions based on your use case.

### 3. Running the Application
To run the application with a specified CSV file, use the following steps:
1. **Build the Project:**
```
go build -o repo-scorer cmd/main.go
```
2. **Run the Application with Flags:**
Use the `-csv` flag to specify the path to your CSV file:
```
./repo-scorer -csv=path/to/commits.csv -user-weighted=true
```
**Flags:**
- `-csv`: Relative path to the CSV file containing commit data.
- `-user-weighted`: Defines whether to use the `basic` or `user-weighted` strategy (optional, default is basic).

This will output the top 10 most active repositories and their activity scores.

### 4. Unit Testing
To run unit tests for the project, use the following command:
```
go test ./services -v
```
This will run all the tests in the `services` folder and display detailed output.

## Example Commit Data (commits.csv)
Here is an example of how the `commits.csv` file should look:
```
timestamp,user,repository,files,additions,deletions
1743865594,user1,repo1,3,10,2
1743865595,user2,repo1,2,5,1
1743865596,user1,repo2,5,20,4
1743865597,user3,repo1,4,15,3
1743865598,,repo3,2,7,0
```
## Notes
- **Data Validation**: The algorithm will validate that the CSV data is correct. Invalid rows will be skipped, and you will be notified via logs.
- **Logging**: Logs will be written to the specified log file (`app.log` by default). Logs are written at various levels (INFO, DEBUG, WARN, ERROR).
- **Strategy Choice**: The strategy (`basic` vs `user-weighted`) can be chosen at runtime. The `user-weighted` strategy will take into account the number of unique users contributing to each repository when calculating the final score.

## Conclusion
This implementation provides a flexible and scalable way to calculate and rank repository activity scores. It includes data validation, logging, and scoring strategies that can be customized based on your the user's needs.

While the current implementation is solid and functional, there are always opportunities for enhancements and improvements. Some of them are:
- **Memory usage optimization**: The algorithm currently loads all commit records into memory. In the case of very large CSV files, this could result in memory issues. Implementing streaming or batch processing could reduce memory footprint for handling larger datasets.
- **Parallel processing**: The repository ranking and commit scoring could be parallelized to improve performance, especially for large datasets. Using Go's concurrency model (goroutines and channels) could speed up the processing.
- **Structured logging**: Rather than just plain text, structured logs in formats like JSON could help integrate this system with centralized logging tools.
- **Additional scoring strategies**: Other strategies beyond the `BasicStrategy` and `UserWeightedStrategy` could be added. For example, an approach that weighs recent commits higher than older ones.
- **GitHub API**: For real-time analysis, we could integrate with GitHub’s API to fetch commit data dynamically, rather than relying exclusively on CSV files.
- **File format support**: Enhance the system to handle different formats beyond CSV, such as JSON or Excel files. This would make the tool more flexible for various use cases.

## Challenge - Top 10 Repositories

The following are the top 10 most active repositories based on the activity score calculated using the user-weighted scoring algorithm:
```
[
  {
    "Name": "repo476",
    "Score": 2415093.2902600914
  },
  {
    "Name": "repo260",
    "Score": 597995.492416038
  },
  {
    "Name": "repo795",
    "Score": 453072.46112738264
  },
  {
    "Name": "repo920",
    "Score": 370293.95083136455
  },
  {
    "Name": "repo518",
    "Score": 322835.74850759626
  },
  {
    "Name": "repo250",
    "Score": 257084.49746293155
  },
  {
    "Name": "repo1143",
    "Score": 218724.304079929
  },
  {
    "Name": "repo161",
    "Score": 217666.26346569046
  },
  {
    "Name": "repo127",
    "Score": 161700.36553149347
  },
  {
    "Name": "repo1185",
    "Score": 157973.80555601302
  }
]
```

### Formula Used to Calculate Activity Score
The activity score for each repository is calculated based on the following formula:
```
score = weight_commits + (weight_files * number_of_files_changed) + (weight_lines * (additions + deletions))
```
Additionally, we apply **unique user weighting** (if enabled via the `-unique-users` flag) to further adjust the score based on the number of unique users involved in each repository. This adjustment is made using the following formula:
```
adjusted_score = base_score * (1 + ln(1 + unique_users))
```