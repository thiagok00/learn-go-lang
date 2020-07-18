package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	r, err := http.Get("http://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("this is an err", err)
	}

	var user1 user
	err = json.Unmarshal(body, &user1)
	fmt.Println(user1)
	fmt.Println("The end.")
}
