package config

type Configuration struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	RecipePuppy struct {
		Url string `yaml:"url"`
	} `yaml:"recipe_puppy"`
	Giphy struct {
		Url string `yaml:"url"`
		ApiKey string `yaml:"api_key"`
	} `yaml:"giphy"`
}

var Config Configuration
