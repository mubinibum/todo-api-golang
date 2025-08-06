package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/mubinibum/todo-api/controllers" // <- sesuaikan nama module GitHub kamu
)

func setupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/todos", controllers.GetTodos)
    r.POST("/todos", controllers.AddTodo)
    r.PUT("/todos/:id", controllers.UpdateTodo)
    r.DELETE("/todos/:id", controllers.DeleteTodo)
    return r
}

func TestGetTodos(t *testing.T) {
    router := gin.Default()
    router.GET("/todos", controllers.GetTodos)

    req, _ := http.NewRequest("GET", "/todos", nil)
    resp := httptest.NewRecorder()

    router.ServeHTTP(resp, req)

    assert.Equal(t, 200, resp.Code)
    assert.Contains(t, resp.Body.String(), "[]")
}

func TestAddTodo(t *testing.T) {
    router := setupRouter()

    todo := map[string]interface{}{
        "title": "Belajar Golang",
        "done":  false,
    }

    body, _ := json.Marshal(todo)
    req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")

    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusCreated, resp.Code)
    assert.Contains(t, resp.Body.String(), "Belajar Golang")
}

func TestUpdateTodo(t *testing.T) {
    router := setupRouter()

    // Tambahkan dulu todo agar bisa di-update
    todo := map[string]interface{}{"title": "Belajar Golang", "done": false}
    body, _ := json.Marshal(todo)
    req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    // Update todo dengan ID 1
    updated := map[string]interface{}{"title": "Sudah Bisa Golang", "done": true}
    updatedBody, _ := json.Marshal(updated)
    req2, _ := http.NewRequest("PUT", "/todos/1", bytes.NewBuffer(updatedBody))
    req2.Header.Set("Content-Type", "application/json")
    resp2 := httptest.NewRecorder()
    router.ServeHTTP(resp2, req2)

    assert.Equal(t, http.StatusOK, resp2.Code)
    assert.Contains(t, resp2.Body.String(), "Sudah Bisa Golang")
}

func TestDeleteTodo(t *testing.T) {
    router := setupRouter()

    // Tambah dulu todo
    todo := map[string]interface{}{"title": "Hapus Saya", "done": false}
    body, _ := json.Marshal(todo)
    req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    // Hapus todo dengan ID 1
    req2, _ := http.NewRequest("DELETE", "/todos/1", nil)
    resp2 := httptest.NewRecorder()
    router.ServeHTTP(resp2, req2)

    assert.Equal(t, http.StatusOK, resp2.Code)
    assert.Contains(t, resp2.Body.String(), "Deleted")
}

