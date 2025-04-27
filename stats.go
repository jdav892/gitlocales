package main 

import (
  "fmt"
)

func stats(email string) {
  commits := processRepositories(email)
  printCommitStats(commits)
}

//processRepositories given a user email
//returns commits made in the last 6 months
func processRepositories(email string) map[int]int {
  filePath := getDotFilePath()
  repos := parseFileLinesToSlice(filePath)
  daysInMap := daysInLastSixMonths

  commits := make(map[int]int, daysInMap)
  for i := daysInMap; i > 0; i-- {
    commits[i] = 0
  }

  for _, path := range repos {
    commits = fileCommits(email, path, commits)
  }
  
  return commits
}

