package main

import (
    "github.com/gin-gonic/gin"
    "github.com/K0725/task-tracker/database"
    "github.com/K0725/task-tracker/routes"
)

func main() {
    r := gin.Default()
    database.ConnectDB()

    routes.RegisterTaskRoutes(r) // 👈 Register your /tasks routes

    r.Run() // :8080
}
