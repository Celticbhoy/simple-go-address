package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Entry struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}

var Entries = []Entry{
	{Id: 1, Firstname: "Steve", Lastname: "Putterson", Email: "ste@putting.ca", Telephone: "226-225-2233"},
	{Id: 2, Firstname: "Sega", Lastname: "Nintendo", Email: "games@ally.jp", Telephone: "522-223-2094"},
}

var EntryId int = 2

// Root Sanity Test
func landingPage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World"})
}

// GET ALL
func allEntries(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Entries)
}

// GET Single
func singleEntry(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	for index, element := range Entries {
		if element.Id == id {
			c.IndentedJSON(http.StatusOK, Entries[index])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Id doesnt exist"})
}

// POST Create New
func addEntry(c *gin.Context) {
	var newEntry Entry
	if err := c.BindJSON(&newEntry); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Entry"})
		return
	}
	newEntry.Id = EntryId + 1
	EntryId += 1
	Entries = append(Entries, newEntry)
}

// DELETE Single
func deleteEntry(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	for index, element := range Entries {
		if element.Id == id {
			Entries = append(Entries[:index], Entries[index+1:]...)
			c.IndentedJSON(http.StatusAccepted, Entries)
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Id doesnt exist"})
}

func main() {
	router := gin.Default()
	router.GET("/address_book", allEntries)
	router.POST("/address_book", addEntry)
	router.GET("/address_book/:id", singleEntry)
	router.DELETE("/address_book/:id", deleteEntry)
	router.GET("/", landingPage)
	router.Run("localhost:8080")
}
