package dirpass

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileEntry struct {
	Path      string // Путь к файлу
	IsDir     bool   // Является ли файл директорией
	IsSymlink bool   // Является ли файл симлинком
	Target    string // Путь, на который указывает симлинк (если это симлинк)
}

func GetFileEntries(path string) ([]FileEntry, error) {
	var entries []FileEntry

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fullPath := filepath.Join(path, file.Name())
		info, err := os.Lstat(fullPath)
		if err != nil {
			return nil, err
		}

		entry := FileEntry{
			Path:     fullPath,
			IsDir:    info.IsDir(),
			IsSymlink: info.Mode()&os.ModeSymlink != 0,
		}

		if entry.IsSymlink {
			target, err := os.Readlink(fullPath)
			if err != nil {
				entry.Target = "[broken]"
			} else {
				_, err := os.Stat(filepath.Join(path, target))
				if err != nil {
					entry.Target = "[broken]"
				} else {
					entry.Target = target
				}
			}
		}

		entries = append(entries, entry)

		if entry.IsDir {
			subEntries, err := GetFileEntries(fullPath)
			if err != nil {
				return nil, err
			}
			entries = append(entries, subEntries...)
		}
	}

	return entries, nil
}
func FilterByType(entries []FileEntry, includeFiles, includeDirs, includeSymlinks bool, ext string) []FileEntry {
	var filtered []FileEntry

	ext = strings.TrimPrefix(ext, ".")

	for _, entry := range entries {
		if includeFiles && !entry.IsDir && !entry.IsSymlink {
			if ext == "" || strings.HasSuffix(entry.Path, "."+ext) {
				filtered = append(filtered, entry)
			}
		} else if includeDirs && entry.IsDir {
			filtered = append(filtered, entry)
		} else if includeSymlinks && entry.IsSymlink {
			filtered = append(filtered, entry)
		} else {
			return entries
		}
	}

	return filtered
}
func PrintFiles(entries []FileEntry) {
	for _, entry := range entries {
		if entry.IsSymlink {
			if entry.Target == "[broken]" {
				fmt.Printf("%s -> [broken]\n", entry.Path)
			} else {
				fmt.Printf("%s -> %s\n", entry.Path, entry.Target)
			}
		} else {
			if entry.IsDir {
				fmt.Printf("%s [DIR]\n", entry.Path)
			} else {
				fmt.Printf("%s [FILE]\n", entry.Path)
			}
		}
	}
}
