package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServer(a *fiber.App) {
	// run server
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Opps.. Server is not ruunning! Reason: %v", err)
	}
}

/*
go fiber에서 우아한 종료 (graceful shutdown)은 fiber내장함수 Shutdown()을 제공해
서버 종료되기 전에 현재 처리 중인 요청들을 모두 완료한 후 서버를 안전하게 종료합니다.
graceful shutdown으로 현제 처리중인 연결들을 마무리하고 클라이언트의 요청에 대해 정상적인 응답을 보낸 후 종료합니다.
*/
func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
