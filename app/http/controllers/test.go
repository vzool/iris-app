package controllers

import (
	"github.com/joho/godotenv"
	"github.com/kataras/iris"
)

// GetPing Handler
func GetPing(ctx iris.Context) {

	env, _ := godotenv.Read()

	ctx.JSON(iris.Map{
		"message": "pong",
		"env":     env,
	})
}
