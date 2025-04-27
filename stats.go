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

//fillCommits given a repository found in `path`, gets the commits and 
//puts them in the `commits` map, returning it when completed
func fillCommits(email string, path string, commits map[int]int) map[int]int {
  //instantiate a git repo object from path
  repo, err := git.PlainOpen(path)
  if err != nil {
    panic(err)
  }
  //get the HEAD reference
  ref, err := repo.Head()
  if err != nil {
    panic(err)
  }
  //get the commits history starting from HEAD
  iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})
  if err != nil {
    panic(err)
  }
  //iterate the commits
  offset := calcOffset()
  err = iterator.ForEach(func(c *object.Commit) error {
    daysAgo := countDaysSinceDate(c.Auther.When) + offset
    
    if c.Author.Email != email {
      return nil
    }

    if daysAgo != outOfRange {
      commits[daysAgo]++
    }

    return nil

  })
  if err != nil {
    panic(err)
  }

  return commits
}

