package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name   string `json:"name"`
	NumPublicRepos int `json:"public_repos"`
}

func main() {
	resp, err := http.Get("https://api.github.com/users/pablomoreno61")
	if err != nil {
		log.Fatalf("error: can't call github.com")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	reader := bytes.NewBufferString(string(data))

	jsonDecoder := json.NewDecoder(reader)

	user := &User{}
	if err := jsonDecoder.Decode(user); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	fmt.Printf("got: %+v\n", user)
}
