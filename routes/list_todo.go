package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"user/db"
	"user/models"

	"github.com/gin-gonic/gin"
)

func ListTodos(c *gin.Context) {
    var todos []models.Todo

    userID := c.Param("user_id")

    pageStr := c.Query("page")
    pageSizeStr := c.Query("pageSize")
    sortDirection := c.Query("sortDirection") // "asc" or "desc"

    defaultPage := 1
    defaultPageSize := 10

    page, err := strconv.Atoi(pageStr)
    if err != nil || page <= 0 {
        page = defaultPage
    }

    pageSize, err := strconv.Atoi(pageSizeStr)
    if err != nil || pageSize <= 0 {
        pageSize = defaultPageSize
    }

    offset := (page - 1) * pageSize
    fmt.Println(offset)

    orderBy := "created"
    if sortDirection == "desc" {
        orderBy = "-" + orderBy
    }

    query := db.Session.Query(`
        SELECT id, user_id, title, description, status, created, updated FROM todos
        WHERE user_id = ? ALLOW FILTERING`, userID)

    iter := query.Iter()

    for {
        var todo models.Todo
        if !iter.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Status, &todo.Created, &todo.Updated) {
            break
        }
        todos = append(todos, todo)
    }

    if err := iter.Close(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, todos)
}
