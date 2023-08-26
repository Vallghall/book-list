package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// WithLogger - middleware decorator that adds given logger
// to handler locals, and info-logs handler execution time
func WithLogger(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("logger", logger)

		start := time.Now()
		logger.Info("handle started",
			zap.String("path", c.Route().Path),
			zap.String("method", c.Route().Method),
			zap.Time("started at", start),
		)

		defer func() {
			logger.Info("handle ended",
				zap.String("path", c.Route().Path),
				zap.String("method", c.Route().Method),
				zap.Int64("finished in s", time.Since(start).Milliseconds()),
			)
		}()

		err := c.Next()
		if err != nil {
			logger.Error(
				"endpoint error",
				zap.String("path", c.Route().Path),
				zap.Error(err),
			)
		}

		return err
	}
}
