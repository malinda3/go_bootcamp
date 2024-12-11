package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "strings"
    "sync"
)

var (
    wg             sync.WaitGroup
    countLinesFlag bool
    countCharsFlag bool
    countWordsFlag bool
)

func init() {
    flag.BoolVar(&countLinesFlag, "l", false, "Count number of lines in each file")
    flag.BoolVar(&countCharsFlag, "m", false, "Count number of characters in each file")
    flag.BoolVar(&countWordsFlag, "w", false, "Count number of words in each file")
}

func countLines(filePath string) int {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatalf("Error while processing %s: %v", filePath, err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lineCount := 0
    for scanner.Scan() {
        lineCount++
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error while scanning %s: %v", filePath, err)
    }
    return lineCount
}

func countCharacters(filePath string) int {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatalf("Error while processing %s: %v", filePath, err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    var charCount int
    for {
        _, err := reader.ReadByte()
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatalf("Error while reading %s: %v", filePath, err)
        }
        charCount++
    }

    return charCount
}

func countWords(filePath string) int {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatalf("Error while processing %s: %v", filePath, err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    wordCount := 0
    for scanner.Scan() {
        text := strings.Fields(scanner.Text())
        wordCount += len(text)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error while counting words in %s: %v", filePath, err)
    }

    return wordCount
}

type Result struct {
    FileName string
    Value    int
}

func processFile(filePath string) Result {
    defer wg.Done()

    var value int
    if countLinesFlag {
        value = countLines(filePath)
    } else if countCharsFlag {
        value = countCharacters(filePath)
    } else if countWordsFlag {
        value = countWords(filePath)
    }

    return Result{FileName: filePath, Value: value}
}

func checkFlags() {
    flagSet := 0
    if countLinesFlag {
        flagSet++
    }
    if countCharsFlag {
        flagSet++
    }
    if countWordsFlag {
        flagSet++
    }

    switch flagSet {
    case 0:
        countWordsFlag = true
    case 1:
    default:
      log.Fatal("You can specify just one of this flags: -l, -m, -w.")
    }
}

func main() {
    flag.Parse()

    files := flag.Args()
    if len(files) == 0 {
        log.Fatal("Please, specify at least1  file")
    }

    checkFlags()

    results := make([]Result, len(files))

    for i, file := range files {
        wg.Add(1)
        go func(index int, filePath string) {
            result := processFile(filePath)
            results[index] = result
        }(i, file)
    }

    wg.Wait()

    for _, result := range results {
        if countLinesFlag {
            fmt.Printf("lines %s: %d\n", result.FileName, result.Value)
        } else if countCharsFlag {
            fmt.Printf("characters %s: %d\n", result.FileName, result.Value)
        } else if countWordsFlag {
            fmt.Printf("words %s: %d\n", result.FileName, result.Value)
        }
    }
}
