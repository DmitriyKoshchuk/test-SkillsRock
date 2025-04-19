package main

import (
	"log"

	"github.com/DmitriyKoshchuk/test-SkillsRock/database"
	"github.com/DmitriyKoshchuk/test-SkillsRock/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Инициализация БД
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.CloseDB()

	app := fiber.New()
	app.Use(logger.New())

	// Маршруты
	app.Get("/tasks", handlers.GetTasks)
	app.Post("/tasks", handlers.CreateTask)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)

	log.Fatal(app.Listen(":3000"))
}
