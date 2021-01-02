package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {
	p1 := person{
		First: "FirsName1",
	}
	p2 := person{
		First: "FirstName2",
	}
	px := []person{p1, p2}
	bs, err := json.Marshal(px)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(bs))
}
