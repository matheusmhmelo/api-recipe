package recipes

import (
)

func GetRecipes(ingredients string) ([]string, error) {
	i, err := format(ingredients)

	if err != nil {
		return nil, err
	}
	return i, nil
}
