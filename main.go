package main

import (
    "github.com/gin-gonic/gin"
    "user/db"
  	"user/routes"
)

func main() {
    r := gin.Default()

    // Initialize ScyllaDB connection
    db.Init()

    // Define routes
    r.POST("/todos", routes.CreateTodo)
    r.GET("/todos/:user_id", routes.ListTodos)
	r.PUT("/todos/:user_id/:post_id",routes.UpdateTodo)
	r.DELETE("/todos/:user_id/:post_id",routes.DeleteTodo)
    // Add other routes for update and delete

    // Run the application
    if err := r.Run(":8080"); err != nil {
        panic(err)
    }
}
