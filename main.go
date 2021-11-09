package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

// func (u URL) String() string {
// 	return "hi"
// }

type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description`
	Payload     string `json:"payload,omitempty"`
}

// func (u URLDescription) String() string {
// 	return "Hello I'm the URL Description"
// }

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
	}
	// fmt.Println(data)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
}

func main() {

	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
