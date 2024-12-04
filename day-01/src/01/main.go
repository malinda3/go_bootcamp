package main

import (
	"fmt"
	"00/pkg/dbreader" 
	"00/pkg/reader"
)
func compare(oldRecipeBook, newRecipeBook *reader.RecipeBook) {
	oldCakes := make(map[string]reader.Cake)
	for _, cake := range oldRecipeBook.Cakes {
		oldCakes[cake.Name] = cake
	}

	newCakes := make(map[string]reader.Cake)
	for _, cake := range newRecipeBook.Cakes {
		newCakes[cake.Name] = cake
	}

	for cakeName := range newCakes {
		if _, exists := oldCakes[cakeName]; !exists {
			fmt.Printf("ADDED cake \"%s\"\n", cakeName)
		}
	}

	for cakeName := range oldCakes {
		if _, exists := newCakes[cakeName]; !exists {
			fmt.Printf("REMOVED cake \"%s\"\n", cakeName)
		}
	}

	for cakeName, oldCake := range oldCakes {
		if newCake, exists := newCakes[cakeName]; exists {
			if oldCake.Time != newCake.Time {
				fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", cakeName, newCake.Time, oldCake.Time)
			}

			oldIngredients := make(map[string]reader.Ingredient)
			for _, ingredient := range oldCake.Ingredients {
				oldIngredients[ingredient.IngredientName] = ingredient
			}

			for _, newIngredient := range newCake.Ingredients {
				oldIngredient, exists := oldIngredients[newIngredient.IngredientName]
				if !exists {
					fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", newIngredient.IngredientName, cakeName)
				} else {
					if oldIngredient.IngredientCount != newIngredient.IngredientCount {
						fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", newIngredient.IngredientName, cakeName, newIngredient.IngredientCount, oldIngredient.IngredientCount)
					}
					if oldIngredient.IngredientUnit != newIngredient.IngredientUnit {
						fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", newIngredient.IngredientName, cakeName, newIngredient.IngredientUnit, oldIngredient.IngredientUnit)
					}
				}
			}

			for _, oldIngredient := range oldCake.Ingredients {
				if _, exists := oldIngredients[oldIngredient.IngredientName]; !exists {
					fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", oldIngredient.IngredientName, cakeName)
				}
			}
		}
	}
}

func main() {
	oldDataBase := dbreader.NewDBReader("../example.xml")
	oldRecipeBook, err := oldDataBase.Run()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	newDataBase := dbreader.NewDBReader("../example.json")
	newRecipeBook, err := newDataBase.Run()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	compare(oldRecipeBook, newRecipeBook)
}
