package main

import (
	"github.com/matheusmhmelo/api-recipe/internal"
	"github.com/matheusmhmelo/api-recipe/internal/config"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
)

func main() {

	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config.Config)
	if err != nil {
		panic(err)
	}

	s := internal.NewServer()
	log.Println("waiting routes...")
	log.Fatal(http.ListenAndServe(config.Config.Server.Port, s.Router))
}
