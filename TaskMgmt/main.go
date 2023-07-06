package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	DueDate     string `json:"due_date" binding:"required"`
	Status      string `json:"status"`
}

func main() {
	InitDB()

	router := gin.Default()

	router.POST("/tasks", CreateTasks)
	router.GET("/tasks/:id", GetTask)
	router.PUT("/tasks/:id", UpdateTask)
	router.DELETE("/tasks/:id", DeleteTask)
	router.GET("/tasks", ListTasks)

	router.Run(":8080")
}

// DATABASE CREATION
func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	// Create the tasks table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		due_date TEXT,
		status TEXT
	);`)
	if err != nil {
		log.Fatal(err)
	}
}

// CREATE
func CreateTasks(c *gin.Context) {
	var tasks []Task
	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTasks := make([]Task, len(tasks))
	for i, task := range tasks {
		// Insert the task into the database
		result, err := db.Exec("INSERT INTO tasks (title, description, due_date, status) VALUES (?, ?, ?, ?)",
			task.Title, task.Description, task.DueDate, task.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Get the ID of the inserted task
		id, _ := result.LastInsertId()
		task.ID = int(id)

		createdTasks[i] = task
	}

	c.JSON(http.StatusOK, createdTasks)
}

// RETRIVE
func GetTask(c *gin.Context) {
	id := c.Param("id")

	var task Task
	err := db.QueryRow("SELECT * FROM tasks WHERE id = ?", id).Scan(
		&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// UPDATE

func UpdateTask(c *gin.Context) {
	id := c.Param("id") // Get the id parameter from the request URL

	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the existing task from the database
	var existingTask Task
	err := db.QueryRow("SELECT * FROM tasks WHERE id = ?", id).Scan(&existingTask.ID, &existingTask.Title, &existingTask.Description, &existingTask.DueDate, &existingTask.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the fields of the existing task
	existingTask.Title = task.Title
	existingTask.Description = task.Description
	existingTask.DueDate = task.DueDate
	existingTask.Status = task.Status

	// Update the task in the database
	_, err = db.Exec("UPDATE tasks SET title = ?, description = ?, due_date = ?, status = ? WHERE id = ?",
		existingTask.Title, existingTask.Description, existingTask.DueDate, existingTask.Status, existingTask.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingTask)
}

// DELETE
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	// Delete the task from the database
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

// LIST_ALL
func ListTasks(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}
