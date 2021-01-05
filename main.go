package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	
	http.HandleFunc("/encode", encodeReq)
	http.HandleFunc("/decode", decodeReq)
	http.ListenAndServe(":8080", nil)
}

func encodeReq(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "FirsName1",
	}
	err := json.NewEncoder(w).Encode(&p1)
	if err != nil {
		log.Fatal(err)
	}
}

func decodeReq(w http.ResponseWriter, r *http.Request) {
	p1 := person{}
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PersonDetails after decoding: ", p1)
}
