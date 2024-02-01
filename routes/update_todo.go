package routes

import (
    "net/http"
    "time"
	"github.com/gocql/gocql"
    "github.com/gin-gonic/gin"
	"user/db"
	"user/models"
 
)


func UpdateTodo(c *gin.Context) {
	var updatedTodo models.Todo
	err := c.BindJSON(&updatedTodo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.Param("user_id")
	todoID := c.Param("post_id")

 
	 
	// Assuming todoID is a string (UUID), modify the type based on your implementation
	// If todoID is an integer, you can use it directly in the query without parsing
	// If it's a different type, adjust accordingly
	query := db.Session.Query(`
		UPDATE todos SET title = ?, description = ?, status = ?, updated = ?
		WHERE user_id = ? AND id = ?`,
		updatedTodo.Title, updatedTodo.Description, updatedTodo.Status, time.Now(), userID, todoID)

	if err := query.Exec(); err != nil {
		if err == gocql.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "TODO item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update TODO item", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "TODO item updated successfully"})
}