package main 

import (
  "flag"
  "fmt"
)


//scan given a path crawls it and its subfolders
//searching for Git repositories
func scan(folder string) {
  fmt.Printf("Found folder:\n\n")
  repositores := recursiveScanFolder(folder)
  filePath := getDotFilePath()
  addNewSliceElementsToFile(filePath, repositories)
  fmt.Printf("\n\nSuccessfully added\n\n")

}
//scanGitFolders returns a list of subfolders of 'folder' ending with '.git'
//Returns the base folder of the repo, the .git folder parent.
//Recursively searches in the subfolders by passing an existing 'folders' slice.
func scanGitFolders(folder []string, folder string) []string {
  
}

//stats generates a nice graph of your Git contributions
func stats(email string) {
  print("stats")
}

func main() {
  var folder string
  var email string
  flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
  flag.StringVar(&email, "email", "your@email.com", "the email to scan")
  flag.Parse()

  if folder != "" {
    scan(folder)
    return
  }

  stats(email)

}
