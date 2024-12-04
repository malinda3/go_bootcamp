package main

import (
	"bufio"
	"fmt"
	"os"
)

func compareFiles(oldFile, newFile string) error {
	oldF, err := os.Open(oldFile)
	if err != nil {
		return fmt.Errorf("failed to open old file: %w", err)
	}
	defer oldF.Close()

	newF, err := os.Open(newFile)
	if err != nil {
		return fmt.Errorf("failed to open new file: %w", err)
	}
	defer newF.Close()

	oldScanner := bufio.NewScanner(oldF)
	newScanner := bufio.NewScanner(newF)

	var oldLine, newLine string
	var oldHasMore, newHasMore bool

	oldHasMore = oldScanner.Scan()
	if oldHasMore {
		oldLine = oldScanner.Text()
	}

	newHasMore = newScanner.Scan()
	if newHasMore {
		newLine = newScanner.Text()
	}

	for oldHasMore || newHasMore {
		if oldHasMore && (!newHasMore || oldLine < newLine) {
			fmt.Printf("REMOVED %s\n", oldLine)
			oldHasMore = oldScanner.Scan()
			if oldHasMore {
				oldLine = oldScanner.Text()
			}
		} else if newHasMore && (!oldHasMore || newLine < oldLine) {
			fmt.Printf("ADDED %s\n", newLine)
			newHasMore = newScanner.Scan()
			if newHasMore {
				newLine = newScanner.Text()
			}
		} else {
			oldHasMore = oldScanner.Scan()
			if oldHasMore {
				oldLine = oldScanner.Text()
			}
			newHasMore = newScanner.Scan()
			if newHasMore {
				newLine = newScanner.Text()
			}
		}
	}

	if err := oldScanner.Err(); err != nil {
		return fmt.Errorf("error reading old file: %w", err)
	}
	if err := newScanner.Err(); err != nil {
		return fmt.Errorf("error reading new file: %w", err)
	}

	return nil
}

func main() {
	oldFile := "snapshot1.txt"
	newFile := "snapshot2.txt"

	err := compareFiles(oldFile, newFile)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
