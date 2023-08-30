package main

import (
	_ "github.com/corgi93/go-fiber-rest-api/docs" // swagger
	"github.com/corgi93/go-fiber-rest-api/pkg/configs"
	"github.com/corgi93/go-fiber-rest-api/pkg/middleware"
	"github.com/corgi93/go-fiber-rest-api/pkg/routes"
	"github.com/corgi93/go-fiber-rest-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // godotenv의 autoload로 .env를 os.Getenv로 읽어올 수 있다.
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	println("go! start")
	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start the Fiber app and listen on port 3000.
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

	utils.StartServerWithGracefulShutdown(app)
}
