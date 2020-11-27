package recipes

import (
	"github.com/matheusmhmelo/api-recipe/internal/config"
	"github.com/matheusmhmelo/api-recipe/internal/domain"
	"github.com/matheusmhmelo/api-recipe/internal/infrastructure/cache"
	"github.com/matheusmhmelo/api-recipe/internal/services/format"
	giphyService "github.com/matheusmhmelo/api-recipe/internal/services/giphy"
	"github.com/matheusmhmelo/api-recipe/internal/services/recipepuppy"
	"log"
	"strings"
)

func GetRecipes(ingredients, page string) (domain.RecipesResponse, error) {
	recipes := domain.RecipesResponse{}

	i, formattedIngredients, err := format.Format(ingredients)
	if err != nil {
		log.Println(err)
		return recipes, err
	}

	foundRecipes, err := searchRecipes(formattedIngredients, page)
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

	rpuppy := recipepuppy.New()
	results, err := rpuppy.Search(ingredients, page)
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

	_cache, err := cache.NewRedis(config.Config.Redis)
	if err != nil {
		log.Println(err)
	}
	if _cache != nil {
		defer _cache.Close()
	}

	for _, recipe := range recipesApi {
		mapRecipe := recipe.(map[string]interface{})
		title := replacer.Replace(mapRecipe["title"].(string))

		modelRecipe := domain.Recipe{
			Title:       title,
			Link:        mapRecipe["href"].(string),
			Ingredients: format.Ingredients(mapRecipe["ingredients"].(string)),
		}

		giphy := giphyService.New()
		gif, err := giphy.Find(title, _cache)
		if err != nil {
			return nil, err
		}

		modelRecipe.Gif = gif

		recipes = append(recipes, modelRecipe)
	}

	return recipes, nil
}