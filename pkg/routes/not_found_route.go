package routes

import "github.com/gofiber/fiber/v2"

// 404에러 라우트
func NotFoundRoute(a *fiber.App) {
	// Register new special route.
	a.Use(
		// 익명함수
		func(c *fiber.Ctx) error {
			// HTTP status 404와  json response 반환
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "sorry, endpoint is not found",
			})
		},
	)
}
