package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"task/internal/logger"
)

// RequestLogger logs all incoming HTTP requests
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Read body (for logging)
		var body []byte
		if c.Method() != fiber.MethodGet {
			b := c.Body()
			body = make([]byte, len(b))
			copy(body, b)
		}

		// Process request
		err := c.Next()

		// Log details
		logger.Logger.Info("HTTP Request",
			logger.ZapString("method", c.Method()),
			logger.ZapString("path", c.Path()),
			logger.ZapString("body", string(body)),
			logger.ZapInt("status", c.Response().StatusCode()),
			logger.ZapDuration("latency", time.Since(start)),
		)

		return err
	}
}
