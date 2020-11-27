package domain

type Recipe struct {
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Link        string   `json:"link"`
	Gif         string   `json:"gif"`
}

type RecipesResponse struct {
	Keywords []string `json:"keywords"`
	Recipes  []Recipe
}
