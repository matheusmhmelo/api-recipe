package recipes

import (
	"github.com/matheusmhmelo/api-recipe/internal/domain"
)

func GetRecipes(ingredients, page string) (domain.RecipesResponse, error) {
	recipes := domain.RecipesResponse{}

	i, err := format(ingredients)
	if err != nil {
		return recipes, err
	}

	foundRecipes, err := searchRecipes(ingredients, page)
	if err != nil {
		return recipes, err
	}

	recipes.Keywords = i
	recipes.Recipes = foundRecipes

	return recipes, nil
}

func searchRecipes(ingredients string, page string) ([]domain.Recipe, error) {
	return nil, nil
}
