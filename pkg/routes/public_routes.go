package routes

import (
	"github.com/corgi93/go-fiber-rest-api/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	// route 그룹 생성
	route := a.Group("/api/v1")
	// server beat check
	route.Get("/beat", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"beat": "ok",
		})
	})

	// GET method의 라우터
	route.Get("/books", controllers.GetBooks)

	// token 생성
	route.Get("/token/new", controllers.GetNewAccessToken)

}
