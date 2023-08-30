package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddleware(a *fiber.App) {
	println("fiber middleware..")
	// fiber app에 미들웨어 적용
	// cors , logger
	a.Use(
		// cors 적용
		cors.New(),
		// 간단한 logger추가
		logger.New(),
	)
}
