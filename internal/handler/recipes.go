package handler

import (
	"encoding/json"
	"fmt"
	"github.com/matheusmhmelo/api-recipe/internal/services/recipes"
	"net/http"
)

//Recipes found recipes with the ingredients
func Recipes(w http.ResponseWriter, r *http.Request) {
	ingredients := r.URL.Query().Get("i")

	if ingredients == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf(`{ "error": "%s" }`, invalidParameters)))
		return
	}

	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	results, err := recipes.GetRecipes(ingredients, page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf(`{ "error": "%s" }`, err.Error())))
		return
	}

	ret, _ := json.Marshal(results)
	_, _ = w.Write(ret)

}
