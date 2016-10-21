package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"go-todo/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate task
		var task models.Task
		// Map JSON body to task
		c.Bind(&task)
		// Add task
		id, err := models.PutTask(db, task.Name)
		// Return success response
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}
	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeleteTask(db, id)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}
