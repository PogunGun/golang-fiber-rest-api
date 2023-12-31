package routes

import (
	"github.com/PogunGun/golang-fiber-rest-api/controllers"
	"github.com/PogunGun/golang-fiber-rest-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthenticated)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.LogOut)

	// CRUD USER
	app.Put("/api/users/info", controllers.UpdateInfo)
	app.Put("/api/users/password", controllers.UpdatePassword)

	app.Get("/api/users", controllers.AllUser)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	// CRUD ORDERS
	app.Get("/api/product", controllers.AllOrders)
	app.Post("/api/product", controllers.Export)
	// CRUD PRODUCTS

	app.Get("/api/product", controllers.AllProducts)
	app.Post("/api/product", controllers.CreateProduct)
	app.Get("/api/product/:id", controllers.GetProduct)
	app.Put("/api/product/:id", controllers.UpdateProduct)
	app.Delete("/api/product/:id", controllers.DeleteProduct)

	// CRUD Role
	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	//Permission
	app.Get("/api/permissions", controllers.AllPermissions)

	//Upload
	app.Post("/api/upload", controllers.Upload)
	app.Static("/api/uploads", "./uploads")
}
