package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) getRecipeByID(c *gin.Context) {
	id := c.Param("id")
	recipeID, _ := strconv.Atoi(id)

	recipe, err := app.DB.Recipe(recipeID)
	if err != nil {
		app.errorLog.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, recipe)
}
