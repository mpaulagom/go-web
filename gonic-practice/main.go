package main

import (
	"encoding/json"
	"fmt"
	"jsonapp/gonic-practice/handler"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int
	Name string
	Pass string
}

func main() {
	usr := User{
		Id:   123,
		Name: "John Doe",
		Pass: "123susuu",
	}

	userjson, err := json.Marshal(usr)
	if err != nil {
		panic(err)
	}
	fmt.Println("User: ", string(userjson))

	//create a router with gin
	router := gin.Default()

	router.GET("/hello-world", handler.HelloWorldGet)

	router.GET("/ping", handler.PingGet)

	router.POST("/saludo", handler.SaludoPost)
	//start the server on port 8081
	router.Run(":8081")
}
