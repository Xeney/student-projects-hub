package handlers

import (
	"net/http"
	"time"

	"github.com/Xeney/student-projects-hub/backend/models"

	"github.com/gin-gonic/gin"
)

var projects = []models.Project{
	{
		ID:          1,
		Title:       "Tetris на C#",
		Description: "Консольная версия игры",
		Author:      "Дима",
		CreatedAt:   time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:          2,
		Title:       "Веб-приложение для заметок",
		Description: "SPA приложение на React с бэкендом на Go",
		Author:      "Дамир",
		CreatedAt:   time.Date(2026, 1, 20, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:          3,
		Title:       "API для системы блогов",
		Description: "REST API с JWT аутентификацией",
		Author:      "Дима",
		CreatedAt:   time.Date(2026, 1, 25, 0, 0, 0, 0, time.UTC),
	},
}

// GetProjects возвращает список всех проектов
// GET /api/projects
func GetProjects(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"data":     projects,
		"count":    len(projects),
		"endpoint": "GET /api/projects",
	})
}
