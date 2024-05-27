package tg_bot

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// server
// get, URL, answer(json)
func main() {
	server := gin.Default()
	server.POST("/videos", func(ctx *gin.Context) { //в context закидывае данные
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(200, gin.H{"message": "Video input is Valid!!"})
		}

	})
	server.Run(":8080")
}
