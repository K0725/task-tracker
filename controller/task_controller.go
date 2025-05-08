package controller

import (
	"github.com/K0725/task-tracker/database"
	"github.com/K0725/task-tracker/controller"
	"github.com/gin-gonic/gin"	
)

func main(){
	// GORM is an orm library which interact with the database 
	r := gin.Default()
	database.ConnectDB()

	r.POST("/tasks", controller.CreateTask)

	r.Run()
}