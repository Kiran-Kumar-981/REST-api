package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id     string `json:"id"`
	Item   string `json:"Item"`
	Packed bool   `json:"Packed"`
}

var Todos = []Todo{
	{Id: "1", Item: "laptop", Packed: false},
	{Id: "2", Item: "mobile", Packed: false},
	{Id: "3", Item: "cloths", Packed: false},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:Id", getTodo)
	router.PATCH("/todos/:Id", todoStatus)
	router.POST("/todos", addTodo)
	http.ListenAndServe(":2020", router)
}

func getTodos(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, Todos)
}

func addTodo(ctx *gin.Context) {
	var newTodo Todo
	if err := ctx.BindJSON(&newTodo); err != nil {
		return
	}
	Todos = append(Todos, newTodo)
	ctx.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(ctx *gin.Context) {
	Id := ctx.Param("Id")
	Todo, err := getTodoById(Id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"id message": "todo not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, Todo)
}

func getTodoById(Id string) (*Todo, error) {
	for i, j := range Todos {
		if j.Id == Id {
			return &Todos[i], nil
		}
	}
	return nil, errors.New("todos not found")
}

func todoStatus(ctx *gin.Context) {
	id := ctx.Param("Id")
	Todo, err := getTodoById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"id message": "todo not found"})
		return
	}
	Todo.Packed = !Todo.Packed
	ctx.IndentedJSON(http.StatusOK, Todo)
}
