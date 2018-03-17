package main

import (
	"fmt"
	"github.com/jordyvandomselaar/dnd5/src/modules/characters"
	"log"
	"net/http"
	"os"
)

func main() {
	characters.Init()

	address := fmt.Sprintf("localhost:%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(address, nil))
}
