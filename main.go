// R. Conner Howell - 2018
// --------------------------------
//	TODO REST API Implemented in Go
// --------------------------------

package main

import (
	"github.com/gin-gonic/gin" // The framework
	"net/http" // Used for http.Status_
	"time" // Used for generating id's
	"strconv" // Used for converting strings to ints
)

// Model for a TODO
type TODO struct {
	Id 			int			`json:"id"`
	Done	 	bool		`json:"done"`
	Text		string	`json:"text"`
}

// In-memory data store
// I chose to use a slice instead of a map
// because there's no built-in map -> slice method
// and O(n) lookup on update and delete is reasonable given
// the small data size
var TODOS []TODO

// Create a new TODO
func create (c *gin.Context) {
	// Get text from url parameter
	text := c.Query("text")
	// Append a new TODO to the list
	TODOS = append(TODOS, TODO{
		Id: time.Now().Nanosecond(),
		Done: false,
		Text: text,
	})
	// Return with success
	c.JSON(http.StatusCreated, "Created")
}

// Return all TODO's
func read (c *gin.Context) {
	// We hold this method to be self-evident
	c.JSON(http.StatusOK, TODOS)
}

// Update a single TODO
func update (c *gin.Context) {
	// Get the TODO id from url parameter
	id, _ := strconv.Atoi(c.Query("id"))
	found := false
	// Lookup the TODO to update: O(n)
	for i := range TODOS {
		if (TODOS[i].Id == id) {
			// Update TODO status
			found = true
			TODOS[i].Done = !TODOS[i].Done
		}
	}
	// Return with success if TODO was found
	if (found) {
		c.JSON(http.StatusOK, "Updated")
	} else {
		c.JSON(http.StatusBadRequest, "Id not found")
	}
}

// Delete a TODO
func delete (c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	found := false
	// Lookup the TODO to delete: O(n)
	for i := range TODOS {
		if (TODOS[i].Id == id) {
			found = true
			// Delete TODO from the slice
			TODOS = append(TODOS[:i], TODOS[i+1:]...)
		}
	}
	// Return with success if TODO was found
	if (found) {
		c.JSON(http.StatusOK, "Deleted")
	} else {
		c.JSON(http.StatusBadRequest, "Id not found")
	}
}

func main() {
	// Initialize gin with default middleware
	r := gin.Default()
	// Attach routes to REST methods and CRUD actions on our TODO's
	r.POST("/", 	create)
	r.GET("/", 		read)
	r.PUT("/", 		update)
	r.DELETE("/", delete)
	// Listen on port 1337
	r.Run(":1337")
}
