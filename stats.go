package main 

import (
  "fmt"
  "github.com/go-git/go-git"
  "github.com/go-git/go-git/plumbing/object"
  "sort"
  "time"
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

//getBeginningOfDay given a time. Time calculates the start time of that day 
func getBeginningOfDay(t time.Time) time.Time {
  year, month, day := t.Date()
  startOfDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
  return startOfDay
}

//countDaysSinceDate counts how many days payssed since the passed `date`
func countDaysSinceDate(date time.Time) int {
  days := 0
  now := getBeginningOfDay(time.Now())
  for date.Before(now) {
    date = date.Add(time.Hour * 24)
    days++
    if days > daysInLastSixMonths {
      return outOfRange
    }
  }
  return days
}

//calcOffset determines and returns the amount of days missing to fillCommits
//the last row of the stats graph
func calcOffset() int {
  var offset := int
  weekday := time.Now().Weekday()

  switch weekday {
  case time.Sunday:
    offset = 7
  case time.Monday:
    offset = 6
  case time.Tuesday:
    offset = 5
  case time.Wednesday:
    offset = 4
  case time.Thursday:
    offset = 3
  case time.Friday:
    offset = 2
  case time.Saturday:
    offset = 1
  }

  return offset
}

//printCommitStats prints the commits stats 
func printCommitStats(commits map[int]int) {
  keys := sortMapIntoSlices(commits)
  cols := buildCols(keys, commits)
  printCells(cols)
}

//sortMapIntoSlices returns a slice of indexes of a map, ordered
func sortMapIntoSlices(m map[int]int) []int {
  //order map
  //To store the keys in slice sorted order
  var keys []int
  for k := range m {
    keys = append(keys, k)
  }
  sort.Ints(keys)
  return keys
}
