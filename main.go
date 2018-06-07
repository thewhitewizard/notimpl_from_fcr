package main //author: fcordier

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ajaxify", func(c *gin.Context) {
		c.JSON(404, gin.H{"msg": "this feature is not implemented"})
	})
	router.Run()
}
