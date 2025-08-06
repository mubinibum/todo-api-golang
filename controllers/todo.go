package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/mubinibum/todo-api/models"
)

var todos []models.Todo
var nextID = 1

func GetTodos(c *gin.Context) {
    todos := []models.Todo{}
    c.JSON(http.StatusOK, todos)
}

func AddTodo(c *gin.Context) {
    var newTodo models.Todo
    if err := c.ShouldBindJSON(&newTodo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newTodo.ID = nextID
    nextID++
    todos = append(todos, newTodo)
    c.JSON(http.StatusCreated, newTodo)
}

func UpdateTodo(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var updated models.Todo
    if err := c.ShouldBindJSON(&updated); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, todo := range todos {
        if todo.ID == id {
            todos[i].Title = updated.Title
            todos[i].Done = updated.Done
            c.JSON(http.StatusOK, todos[i])
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func DeleteTodo(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    for i, todo := range todos {
        if todo.ID == id {
            todos = append(todos[:i], todos[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
