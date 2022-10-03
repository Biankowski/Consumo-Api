package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userid"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	api := "https://jsonplaceholder.typicode.com/todos"

	response, _ := http.Get(api)

	responseData, _ := ioutil.ReadAll(response.Body)

	var todos []Todo

	json.Unmarshal([]byte(responseData), &todos)

	for _, todo := range todos {
		fmt.Println("|---------------------------------------------------------------------------------------------------------|")
		fmt.Printf("User Id: %d, Id: %d, Title: %s, Completed: %t \n", todo.UserId, todo.Id, todo.Title, todo.Completed)

	}

}
