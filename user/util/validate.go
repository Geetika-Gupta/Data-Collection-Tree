package util

import (
	"github.com/gin-gonic/gin"
)

//VInsert : will validate insert web request
func VInsert(context *gin.Context) (req model.WebRequest, err error) {
	err := context.ShouldBindJSON(&req)
	return
}

//VGet : will validate get web request
func VGet(context *gin.Context) (req model.Dimension, err error) {
	err := context.ShouldBindJSON(&req)
	return
}