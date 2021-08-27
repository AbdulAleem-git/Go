package main

import (
	"encoding/json"
	"net/http"
)

type Animal struct {
    Name  string `json:"Name"`
    Type  string `json:"Type"`
}

func AnimalsHandler(w http.ResponseWriter, r *http.Request) {
   animals := []Animal{
       Animal{"Alice","Cat"},
       Animal{"Bob","Cat"},
       Animal{"Trinity","Dog"},
   }

   json.NewEncoder(w).Encode(animals)
}


func main() {
    http.HandleFunc("/animals", AnimalsHandler)
    http.ListenAndServe(":1234", nil)
}