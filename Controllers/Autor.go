package Controllers

import (
	"../ApiHelpers"
	"../Models"
	"github.com/gin-gonic/gin"
)

func ListAuthor(c *gin.Context) {
	var author []Models.Author
	
	err := Models.GetAllAuthor(&author)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, author)
	} else {
		ApiHelpers.RespondJSON(c, 200, author)
	}
}