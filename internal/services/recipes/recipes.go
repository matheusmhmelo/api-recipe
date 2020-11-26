package recipes

import (
	"github.com/matheusmhmelo/api-recipe/internal/config"
	"github.com/matheusmhmelo/api-recipe/internal/domain"
	"github.com/matheusmhmelo/api-recipe/internal/utils"
	"log"
	"strings"
)

func GetRecipes(ingredients, page string) (domain.RecipesResponse, error) {
	recipes := domain.RecipesResponse{}

	i, err := format(ingredients)
	if err != nil {
		log.Println(err)
		return recipes, err
	}

	foundRecipes, err := searchRecipes(ingredients, page)
	if err != nil {
		log.Println(err)
		return recipes, err
	}

	recipes.Keywords = i
	recipes.Recipes = foundRecipes

	return recipes, nil
}

func searchRecipes(ingredients string, page string) ([]domain.Recipe, error) {
	var recipes []domain.Recipe

	results, err := search(ingredients, page)
	if err != nil {
		return recipes, err
	}
	if len(results) == 0 {
		return recipes, nil
	}

	recipes, err = handleRecipes(results)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func handleRecipes(recipesApi []interface{}) ([]domain.Recipe, error) {
	var recipes []domain.Recipe
	replacer := strings.NewReplacer("\n", "", "\r", "", "\t", "")

	cache, err := utils.NewRedis(config.Config.Redis)
	if err != nil {
		log.Println(err)
	}
	defer cache.Close()

	for _, recipe := range recipesApi {
		mapRecipe := recipe.(map[string]interface{})
		title := replacer.Replace(mapRecipe["title"].(string))

		modelRecipe := domain.Recipe{
			Title: title,
			Link: mapRecipe["href"].(string),
			Ingredients: formatIngredients(mapRecipe["ingredients"].(string)),
		}

		gif, err := findGif(title, cache)
		if err != nil {
			return nil, err
		}

		modelRecipe.Gif = gif

		recipes = append(recipes, modelRecipe)
	}

	return recipes, nil
}