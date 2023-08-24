package main

import (
	"back_product/auth"
	"back_product/product"
	"back_product/user"

	"github.com/gofiber/fiber/v2"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")

}

func setupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello my api")
	})

	app.Get("/status", status)

	api := app.Group("/api")

	apiLogin := api.Group("/auth")
	apiLogin.Get("/login", auth.LoginHandler)

	apiProduct := api.Group("/product", auth.AuthorizationMiddleware)
	apiProduct.Get("/products", product.GetAllProduct)

	apiUser := api.Group("/user")
	apiUser.Get("/create", user.AddUser)
}
