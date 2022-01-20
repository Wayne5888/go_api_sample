package main

import	(
	"unittest/controller"
	"fmt"
	// "net/http"
	"github.com/gin-gonic/gin" 
	//"database/sql"
	
)

func sqlSample(){

}

func main(){
	server := gin.Default()
	fmt.Println("hello world")

	server.GET("/ping", ping)
	server.GET("/getUserByID", controller.GetUserByID)
	server.GET("/getUserList", controller.GetUserList)
	
	server.Run(":8888")

}

func ping(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong2",
	})
}