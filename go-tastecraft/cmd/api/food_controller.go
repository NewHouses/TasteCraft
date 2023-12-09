package main

import (
	e "tastecraft/internal/entities"

	"github.com/gin-gonic/gin"
)

func (app *application) addFood(c *gin.Context) {
	var food e.Food

	c.BindJSON(&food)

	if len(food.Name) == 0 {
		app.sendBadRequest(c, "food request is not well constructed")
		return
	}

	if err := app.DB.InsertFood(&food); err != nil {
		app.sendConflict(c, err.Error())
		return
	}

	app.sendOK(c)
}
