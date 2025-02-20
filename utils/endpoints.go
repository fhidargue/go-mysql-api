package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"mysql-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetStudentsAPI(c *gin.Context) {
	students, err := GetStudentsDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	if len(students) == 0 {
		c.JSON(http.StatusOK, []models.Student{})
		return
	}
	
	c.JSON(http.StatusOK, students)
}

func GetStudentByIdAPI(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Id param error": err.Error()})
	}

	student, err := GetStudentByIdDB(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, student)
}

func PostStudentAPI(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		var ve validator.ValidationErrors

		if ok := errors.As(err, &ve); ok {
			var missingFields []string

			for _, fe := range ve {
				missingFields = append(missingFields, fe.Field())
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"error":           "Missing required fields",
				"Missing fields":  missingFields,
			})

			return
		}
	}

	id, err := AddStudentDB(student)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add student"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Student added successfully", "Student id": id})
}

func PatchStudentAPI(c *gin.Context) {
	var student models.StudentPatch

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}
	
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	if err := UpdateStudentDB(id, student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Student with ID: %v, was successfully updated", id)})
}

func DeleteStudentAPI(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	_, err = GetStudentByIdDB(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Student with ID %d not found", id)})
		return
	}

	_, err = DeleteStudentDB(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Student with ID %d deleted successfully", id)})
}