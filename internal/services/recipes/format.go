package recipes

import (
	"log"
	"regexp"
	"strings"
)

func format(ingredients string) ([]string, error) {
	i := strings.Split(ingredients, ",")

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
