package dbreader

import (
	"fmt"
	"00/pkg/factory"
	"00/pkg/reader"
)

type DBReader struct {
	FilePath string
}

func NewDBReader(filePath string) *DBReader {
	return &DBReader{FilePath: filePath}
}

func (d *DBReader) Run() (*reader.RecipeBook, error) {
	factory := &factory.ReaderFactory{}
	dbReader, err := factory.GetReader(d.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get reader: %w", err)
	}

	recipeBook, err := dbReader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return recipeBook, nil
}
