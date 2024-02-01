package routes

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
   "user/models"
   "user/db"
)

func CreateTodo(c *gin.Context) {
    var newTodo models.Todo
    if err := c.ShouldBindJSON(&newTodo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newTodo.ID = time.Now().String()
    newTodo.Created = time.Now()
    newTodo.Updated = time.Now()

    if err := db.Session.Query(`
        INSERT INTO todos (id, user_id, title, description, status, created, updated)
        VALUES (?, ?, ?, ?, ?, ?, ?)`,
        newTodo.ID, newTodo.UserID, newTodo.Title, newTodo.Description, newTodo.Status, newTodo.Created, newTodo.Updated).Exec(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newTodo)
}
