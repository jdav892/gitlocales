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
  // trim the last '/'
  folder = strings.TrimSuffix(folder, "/")

  f, err := os.Open(folder)
  if err != nil {
    log.Fatal(err)
  }
  files, err := f.Readdir(-1)
  f.Close()
  if err != nil {
    log.Fatal(err)
  }

  var path string

  for _, file := range files {
    if file.IsDir() {
      path = folder + "/" + file.Name()
      if file.Name() == ".git" {
        path = strings.TrimSuffix(path, "/.git")
        fmt.Println(path)
        folders = append(folders, path)
        continue
      }
      if file.Name() == "vendor" || file.Name() == "node_modules" {
        continue
      }
      folders = scanGitFolders(folders, path)
    }
  }
  return folders
}

//recursiveScanFolder starts the recursive search of git repositories
//living in the 'folder' subtree
func recursiveScanFolder(folder string) []string {
  return scanGitFolders(make([]string, 0), folder)
}

//getDotFilePath returns the dot file for the repos list
//Creates it and the enclosing folder if it does not exist
func getDotFilePath() string {
  usr, err := user.Current()
  if err != nil {
    log.Fatal(err)
  }

  dotFile := usr.HomeDir + "/.gogitlocalstats"

  return dotFile
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
