package controller

import "unittest/func/test"
import "github.com/gin-gonic/gin" 
import "net/http"


func GetUserByID(c *gin.Context){
	id := c.Query("id")
	data := test.GetUserById(id);

	c.JSON(http.StatusOK,data)
}

func GetUserList(c *gin.Context){
	data := test.GetUserList();

	c.JSON(http.StatusOK,data)
}


