package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// this data is different from JSON format. JSON is used for communication between client and the server in REST API.
var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

// function to convert the todo array into JSON and return
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// func displayRoot(context *gin.Context) {
// 	context.IndentedJSON(http.StatusOK, "Hello World\n")
// }

// receive client-data in JSON and convert it to `todo` datatype
func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

// first we need to extract the 'id' from the POST request URL which is there in `context`
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todos not found!")
}

func main() {
	router := gin.Default()
	// router.GET("/", displayRoot)
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo) // the ':' tells that field is dynamic and is called as 'id' here
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")

}
