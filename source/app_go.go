package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Wake up Neo!\nThis is Go App!\nMade with 💖 by Antosha")
	})

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "v2.0")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}