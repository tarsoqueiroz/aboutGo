package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPong(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message:": "Pong!!!"})
}
