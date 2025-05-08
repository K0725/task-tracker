package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/K0725/task-tracker/controller" // ✅ Import the controllers
)

func RegisterTaskRoutes(router *gin.Engine) {
    taskGroup := router.Group("/tasks")
    {
        taskGroup.POST("/", controller.CreateTask)
        // taskGroup.GET("/", controllers.GetTasks)
        // taskGroup.PUT("/:id", controllers.UpdateTask)
        // taskGroup.DELETE("/:id", controllers.DeleteTask)
    }
}
