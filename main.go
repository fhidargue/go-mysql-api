package main

import (
	"mysql-backend/utils"
    "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	utils.ConnectDB()

	router.GET("/students", utils.GetStudentsAPI)
	router.GET("/students/:id", utils.GetStudentByIdAPI)
	router.PATCH("/students/:id", utils.PatchStudentAPI)
	router.DELETE("/students/:id", utils.DeleteStudentAPI)
	router.POST("/students", utils.PostStudentAPI)
	
	router.Run("localhost:8080")
}