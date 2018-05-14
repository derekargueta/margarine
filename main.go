package main

import (
  "flag"
  "fmt"
  "log"
  "math/rand"
  "os"
  "path/filepath"
  "strings"
  "time"
)

/**
 * Creates an empty file in random places around the filesystem. Pretty useless
 * program whose real purpose is for testing file-finding software. Prints out
 * what files it wrote to stdout so you can clean them up :)
 */

const (
  PublicWriteMode = 1 << 2
)

var (
  fileName = flag.String("f", "derp.txt", "name of the file to touch")
  probability = flag.Int("p", 0, "probability of creating the file in a random directory. Out of 100.")
  root = flag.String("r", "/", "starting directory of the spread")
  includeDot = flag.Bool("d", false, "Whether to include dot-files (default is off)")
)

func visit(path string, f os.FileInfo, err error) error {
  if !f.IsDir() {
    return nil
  }

  if !*includeDot && strings.Contains(path, "/.") {
    // skip dot-files if desired
    return nil
  }

  mode := f.Mode()
  if mode & PublicWriteMode != 0 {
    randVal := rand.Intn(1000)
    if randVal < *probability {
      filePath := fmt.Sprintf("%s/%s", path, *fileName)
      _, err := os.Create(filePath)
      if err != nil {
        log.Printf("Failed to write file %s with error %v\n", filePath, err)
      } else {
        fmt.Println(filePath)
      }
    }
  }
  return nil
}

func main() {
  rand.Seed(time.Now().Unix())
  flag.Parse()
  
  if probability == nil {
    log.Fatal("Please provide a -p value")
  }

  err := filepath.Walk(*root, visit)
  if err != nil {
    log.Fatalf("filepath.Walk() returned %v\n", err)
  }

  os.Exit(0)
}