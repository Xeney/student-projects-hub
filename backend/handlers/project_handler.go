package handlers

import (
	"net/http"
	"time"

	"github.com/Xeney/student-projects-hub/backend/models"
	"github.com/gin-gonic/gin"
)

// Временное хранилище в памяти
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
	// Возвращаем JSON массив проектов со статусом 200 OK
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"data":     projects,
		"count":    len(projects),
		"endpoint": "GET /api/projects",
	})
}

// CreateProject создает новый проект
// POST /api/projects
func CreateProject(c *gin.Context) {
	var request models.CreateProjectRequest

	// Привязываем JSON к структуре
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Неверный формат запроса",
			"error":   err.Error(),
		})
		return
	}

	// Базовая валидация
	if err := request.Validate(); err != nil {
		if validationErr, ok := err.(*models.ValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Ошибка валидации",
				"error":   validationErr.Message,
				"field":   validationErr.Field,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Ошибка валидации",
				"error":   err.Error(),
			})
		}
		return
	}

	// Генерируем новый ID (последний ID + 1)
	newID := 1
	if len(projects) > 0 {
		// Находим максимальный ID
		maxID := 0
		for _, project := range projects {
			if project.ID > maxID {
				maxID = project.ID
			}
		}
		newID = maxID + 1
	}

	// Создаем новый проект
	newProject := models.Project{
		ID:          newID,
		Title:       request.Title,
		Description: request.Description,
		Author:      request.Author,
		CreatedAt:   time.Now().UTC(),
	}

	// Добавляем в хранилище
	projects = append(projects, newProject)

	// Возвращаем созданный проект со статусом 201 Created
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Проект успешно создан",
		"data":    newProject,
	})
}
