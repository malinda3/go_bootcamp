package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"io"
)

func main() {
	archiveDir := flag.String("a", "", "Directory to save archives (default: same as log file)")
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("Usage: ./myRotate [-a archive_dir] log_file [log_file ...]")
		os.Exit(1)
	}

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()

			err := archiveLogFile(file, *archiveDir)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error archiving file %s: %v\n", file, err)
			}
		}(file)
	}

	wg.Wait()
}

func archiveLogFile(logFile, archiveDir string) error {
	info, err := os.Stat(logFile)
	if err != nil {
		return fmt.Errorf("failed to stat file: %w", err)
	}

	if info.IsDir() {
		return fmt.Errorf("specified log file is a directory")
	}

	baseName := filepath.Base(logFile)
	formattedTime := info.ModTime().Format("20060102_150405")
	archiveName := fmt.Sprintf("%s_%s.tar.gz", baseName, formattedTime)

	if archiveDir != "" {
		archiveName = filepath.Join(archiveDir, archiveName)
	} else {
		archiveName = filepath.Join(filepath.Dir(logFile), archiveName)
	}

	archiveFile, err := os.Create(archiveName)
	if err != nil {
		return fmt.Errorf("failed to create archive file: %w", err)
	}
	defer archiveFile.Close()

	gzipWriter := gzip.NewWriter(archiveFile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	err = addFileToTar(logFile, tarWriter)
	if err != nil {
		return fmt.Errorf("failed to add file to tar: %w", err)
	}
	os.Remove(logFile)
	fmt.Printf("Archived %s to %s\n", logFile, archiveName)
	return nil
}

func addFileToTar(filePath string, tarWriter *tar.Writer) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat file: %w", err)
	}

	header := &tar.Header{
		Name: filepath.Base(filePath),
		Size: info.Size(),
		Mode: int64(info.Mode()),
		ModTime: info.ModTime(),
	}

	if err := tarWriter.WriteHeader(header); err != nil {
		return fmt.Errorf("failed to write tar header: %w", err)
	}

	_, err = io.Copy(tarWriter, file)
	if err != nil {
		return fmt.Errorf("failed to copy file data to tar: %w", err)
	}

	return nil
}
