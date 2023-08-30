package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SwaggerRoute(a *fiber.App) {
	// route 그룹 생성
	route := a.Group("/swagger")
	// get one user by ID
	route.Get("*", swagger.HandlerDefault)
}
