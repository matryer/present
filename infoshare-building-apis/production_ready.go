package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Msg struct {
	Message string `json:"msg"`
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(Msg{Message: "Hello infoShare"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
