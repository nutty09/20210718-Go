package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type User struct {
	Id   int    `json:"id"` //json literal
	Name string `json:"name"`
}

func main() {
	//struct to json
	u := User{Id: 1, Name: "Somkiat"}
	e := json.NewEncoder(os.Stdout)
	e.Encode(u)

	//json to struct
	const jsonStream = `{"id": 1, "name": "Knock Knock."}`

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m User
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v: %v\n", m.Id, m.Name)
	}
}
