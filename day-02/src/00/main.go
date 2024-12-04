package main

import (
	"00/dirpass"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	showFiles := flag.Bool("f", false, "Show files")
	showDirs := flag.Bool("d", false, "Show directories")
	showSymlinks := flag.Bool("sl", false, "Show symlinks")
	showFilesExt := flag.String("ext", "", "!!WORKS ONLY WITH -f!! Show only files with a certain extension")

	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("Usage: go run [-f(possible with -ext '.extension') -d -sl ] /path/to/dir")
		os.Exit(1)
	}

	path := flag.Arg(0)

	entries, err := dirpass.GetFileEntries(path)
	if err != nil {
		log.Fatal(err)
	}
	filtered := dirpass.FilterByType(entries, *showFiles, *showDirs, *showSymlinks, *showFilesExt)

	dirpass.PrintFiles(filtered)
}
