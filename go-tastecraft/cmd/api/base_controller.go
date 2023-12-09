package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

func (app *application) sendOK(c *gin.Context) {
	app.sendOKWithResult(c, nil)
}

func (app *application) sendOKWithResult(c *gin.Context, result interface{}) {
	var payload response

	payload.Error = false
	payload.Message = "Operation completed successfully"
	payload.Content = result

	app.infoLog.Println("Response OK")
	c.IndentedJSON(http.StatusOK, payload)
}

func (app *application) sendFailure(c *gin.Context, httpStatus int, errorMessage string) {
	var payload response

	payload.Error = true
	payload.Message = errorMessage

	app.errorLog.Println(errorMessage)
	c.IndentedJSON(httpStatus, payload)
}

func (app *application) sendNotFound(c *gin.Context, errorMessage string) {
	app.sendFailure(c, http.StatusNotFound, errorMessage)
}

func (app *application) sendBadRequest(c *gin.Context, errorMessage string) {
	app.sendFailure(c, http.StatusBadRequest, errorMessage)
}

func (app *application) sendConflict(c *gin.Context, errorMessage string) {
	app.sendFailure(c, http.StatusConflict, errorMessage)
}
