package characters

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller struct {
}

func (c Controller) Index(w http.ResponseWriter, r *http.Request) {
	characters := []Character{
		{
			Id:   1,
			Name: "Gentoo",
		},
		{
			Id:   2,
			Name: "Flux",
		},
	}

	data, _ := json.Marshal(characters)

	fmt.Fprintf(w, string(data))
}

type Character struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func Init() {
	c := Controller{}

	http.HandleFunc("/characters", c.Index)
}
