package recipes

import (
	"errors"
	"log"
	"regexp"
	"strings"
)

const ingredientLimitError = "ingredient limit reached, choose only 3"

func format(ingredients string) ([]string, error) {
	i := strings.Split(ingredients, ",")
	if len(i) > 3 {
		err := errors.New(ingredientLimitError)
		return nil, err
	}

	reg, err := regexp.Compile("[^0-9 a-záéíóúàèìòùâêîôûãõç]")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for index, ingredient := range i {
		processedIngredient := reg.ReplaceAllString(ingredient, "")
		i[index] = processedIngredient
	}

	return i, nil
}
