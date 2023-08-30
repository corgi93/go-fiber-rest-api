package routes

import (
	"github.com/corgi93/go-fiber-rest-api/app/controllers"
	"github.com/corgi93/go-fiber-rest-api/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	// POST method
	// 책 생성
	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook)
}
