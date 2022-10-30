package pooling

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	r.POST("/", ProcessCloudStackRequest)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func ProcessCloudStackRequest(c *gin.Context) {
	c.JSON(200, "Oi")
}
