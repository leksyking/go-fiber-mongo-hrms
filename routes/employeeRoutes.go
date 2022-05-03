package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leksyking/go-fiber-mongo-hrms/controllers"
)

var EmployeeRoutes = func(app *fiber.App) {
	app.Get("/api/v1/employee/", controllers.GetAllEmployees)
	app.Post("/api/v1/employee/", controllers.CreateEmployee)
	app.Put("/api/v1/employee/:id", controllers.UpdateEmployee)
	app.Delete("/api/v1/employee/:id", controllers.DeleteEmployee)
}
