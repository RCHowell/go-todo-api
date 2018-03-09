// R. Conner Howell - 2018
// --------------------------------
// TODO REST API Implemented in Go
// --------------------------------

package main

import (
	"net/http" // Used for http.Status_
	"strconv"
	"time"
	// Used for generating id's
	// Used for converting strings to ints

	"github.com/gin-gonic/gin" // The framework
)

// Define how a toDo should be represented
type toDo struct {
	ID   int64  `json:"id"`
	Done bool   `json:"done"`
	Text string `json:"text"`
}

// ToDos is our in memory data store and using a map
// gives us O(1) for Create,Update, and Delete
var toDos map[int64]*toDo

// Create a new toDo
func createToDo(c *gin.Context) {
	// Get text from url parameter
	text := c.Query("text")
	// Use nanoseconds so that toDo.ID is a ~hopefully~ unique value
	id := time.Now().UnixNano()
	// Add the address of our new toDo to the data store
	toDos[id] = &toDo{
		ID:   id,
		Done: false,
		Text: text,
	}
	// Return with success
	c.JSON(http.StatusCreated, "Created")
}

// Respond with a slice of toDos
// Constructing the slice runs in O(n)
func readToDos(c *gin.Context) {
	toDosSlice := make([]toDo, len(toDos))
	i := 0
	for _, val := range toDos {
		toDosSlice[i] = *val
		i++
	}
	c.JSON(http.StatusOK, toDosSlice)
}

// Update a single toDo
func updateToDo(c *gin.Context) {
	// Get the TODO id from url parameter
	id, _ := strconv.Atoi(c.Query("id"))
	toDos[int64(id)].Done = !toDos[int64(id)].Done
	c.JSON(http.StatusOK, "Updated")
}

// Delete a toDo
func deleteToDo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	// Delete takes no action if the key is not present
	delete(toDos, int64(id))
	c.JSON(http.StatusOK, "Deleted")
}

func main() {
	// Initialize gin with default middleware
	r := gin.Default()
	// Initialize our data store
	toDos = make(map[int64]*toDo)
	// Attach routes to REST methods and CRUD actions on our ToDo's
	r.POST("/", createToDo)
	r.GET("/", readToDos)
	r.PUT("/", updateToDo)
	r.DELETE("/", deleteToDo)
	// Listen on port 1337
	r.Run(":1337")
}
