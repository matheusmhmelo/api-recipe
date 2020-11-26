package handler

import (
	"encoding/json"
	"errors"
	"github.com/matheusmhmelo/api-recipe/internal/services/recipes"
	"github.com/matheusmhmelo/api-recipe/internal/utils"
	"net/http"
)

//Recipes found recipes with the ingredients
func Recipes(w http.ResponseWriter, r *http.Request) {
	ingredients := r.URL.Query().Get("i")

	if ingredients == "" {
		err := errors.New("invalid ingredients parameters")
		utils.CreateBadRequestResponse(w, err)
	}

	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	results, err := recipes.GetRecipes(ingredients, page)
	if err != nil {
		utils.CreateBadRequestResponse(w, err)
	}

	ret, _ := json.Marshal(results)
	_, _ = w.Write(ret)

}
