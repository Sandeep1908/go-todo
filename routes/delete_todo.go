package routes

import(
	"net/http"
 
	"github.com/gocql/gocql"
    "github.com/gin-gonic/gin"
 
   "user/db"
)

func DeleteTodo(c *gin.Context) {
	userID := c.Param("user_id")
	todoID := c.Param("post_id")

	
	 

	// Assuming todoID is a string, modify the type based on your implementation
	// If todoID is an integer, you can use it directly in the query without parsing
	// If it's a UUID, you can use gocql.ParseUUID(todoID) to get a UUID object
	// Update the data model and the query accordingly
	// The code below assumes todoID is a string
	query := db.Session.Query(`DELETE FROM todos WHERE user_id = ? AND id = ?`, userID, todoID)

	if err := query.Exec(); err != nil {
		if err == gocql.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "TODO item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete TODO item", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "TODO item deleted successfully"})
}