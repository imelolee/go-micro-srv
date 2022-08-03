package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test1(ctx *gin.Context) {
	fmt.Println("Test1...")
}

func Test2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Test2...")
	}
}

func main() {
	router := gin.Default()

	// 使用中间件
	router.Use(Test1)
	router.Use(Test2())
	router.GET("/", func(context *gin.Context) {
		fmt.Println("GET...")
		context.Writer.WriteString("Hello World!")
	})
	router.Run(":8081")
}
