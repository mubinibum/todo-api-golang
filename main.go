package main

import (
    "github.com/gin-gonic/gin"
    "github.com/mubinibum/todo-api/controllers"
)

func main() {
    r := gin.Default()

    r.GET("/todos", controllers.GetTodos)
    r.POST("/todos", controllers.AddTodo)
    r.PUT("/todos/:id", controllers.UpdateTodo)
    r.DELETE("/todos/:id", controllers.DeleteTodo)

    r.Run(":8080")
}
