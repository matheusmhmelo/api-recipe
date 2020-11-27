package format

import (
	"errors"
	"log"
	"regexp"
	"sort"
	"strings"
)

func Format(ingredients string) ([]string, string, error) {
	i := strings.Split(ingredients, ",")
	if len(i) > 3 {
		err := errors.New("ingredient limit reached, choose only 3")
		return nil, "", err
	}

	reg, err := regexp.Compile("[^0-9 a-záéíóúàèìòùâêîôûãõç]")
	if err != nil {
		log.Fatal(err)
		return nil, "", err
	}

	var formattedIngredients string
	for index, ingredient := range i {
		processedIngredient := reg.ReplaceAllString(ingredient, "")
		i[index] = processedIngredient

		formattedIngredients += processedIngredient + ","
	}

	return i, formattedIngredients[:len(formattedIngredients)-1], nil
}

func Ingredients(ingredients string) []string {
	var i []string
	i = strings.Split(ingredients, ", ")
	sort.Strings(i)

	return i
}
