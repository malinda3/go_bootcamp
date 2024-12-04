package factory

import (
	"fmt"
	"00/pkg/reader"
)

type ReaderFactory struct{}

func (f *ReaderFactory) GetReader(filePath string) (reader.Reader, error) {
	ext := reader.GetFileExtension(filePath)
	switch ext {
	case "json":
		return &reader.JSONReader{FilePath: filePath}, nil
	case "xml":
		return &reader.XMLReader{FilePath: filePath}, nil
	default:
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}
}
