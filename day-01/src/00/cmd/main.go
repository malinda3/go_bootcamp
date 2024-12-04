package main

import (
	"flag"
	"fmt"
	"00/pkg/dbreader" 
)

func main() {
	filePath := flag.String("f", "", "Path to the recipe file (e.g., example.xml or example.json)")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("Error: File path is required.")
		return
	}

	dbReader := dbreader.NewDBReader(*filePath)

	recipeBook, err := dbReader.Run()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Recipe Book:", recipeBook)
}

//go run ./cmd -f ..\example.json