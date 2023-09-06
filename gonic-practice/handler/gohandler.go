package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type SimpleUser struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

func PingGet(c *gin.Context) {

	c.JSON(200, gin.H{
		"msg": "pong",
	})
}
func HelloWorldGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "hellos wolrd!",
	})
}

func SaludoPost(ctx *gin.Context) {
	var usr SimpleUser
	if err := ctx.BindJSON(&usr); err != nil {
		fmt.Println("me vino!: ", usr.LastName)
		ctx.JSON(200, usr)
	}
	ctx.JSON(200, usr)
}
