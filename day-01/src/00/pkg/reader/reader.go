package reader

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
)

type Reader interface {
	Read() (*RecipeBook, error)
}

type Ingredient struct {
	IngredientName  string `json:"ingredient_name" xml:"itemname"`
	IngredientCount string `json:"ingredient_count" xml:"itemcount"`
	IngredientUnit  string `json:"ingredient_unit" xml:"itemunit"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type RecipeBook struct {
	Cakes []Cake `json:"cake" xml:"cake"`
}

type JSONReader struct {
	FilePath string
}

func (j *JSONReader) Read() (*RecipeBook, error) {
	jsonFile, err := os.Open(j.FilePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var recipeBook RecipeBook
	err = json.Unmarshal(byteValue, &recipeBook)
	if err != nil {
		return nil, err
	}
	return &recipeBook, nil
}

type XMLReader struct {
	FilePath string
}

func (x *XMLReader) Read() (*RecipeBook, error) {
	xmlFile, err := os.Open(x.FilePath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	var recipeBook RecipeBook
	err = xml.Unmarshal(byteValue, &recipeBook)
	if err != nil {
		return nil, err
	}
	return &recipeBook, nil
}

func GetFileExtension(fileName string) string {
	return filepath.Ext(fileName)[1:]
}
