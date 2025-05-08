package main

import (
	"github.com/K0725/task-tracker/database"
	"github.com/K0725/task-tracker/controllers"
	"github.com/gin-gonic/gin"	
)

func main(){
	// GORM is an orm library which interact with the database 
	r := gin.Default()
	database.ConnectDB()

	r.POST("/tasks", controllers.CreateTask)

	r.Run()
}