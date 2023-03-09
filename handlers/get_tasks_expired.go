package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTasksExpired(c *gin.Context) {

	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable")

	rows, err := db.Query("SELECT id, title, description, due_date, completed FROM tasks WHERE due_date<$1 ORDER BY id", time.Now())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}
