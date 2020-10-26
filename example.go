package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(func01)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(500, gin.H{
			"message": c.Errors,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func func01(ctx *gin.Context) {
	ctx.Error(errors.New("func01"))
	func02(ctx)
}

func func02(ctx *gin.Context) error {
	err := errors.New("func02")
	if err != nil {
		ctx.Error(err).SetMeta("foo")
		return err
	}
	return err
}
