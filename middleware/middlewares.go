package middleware

// for finding request id ///

import (
	"log"
	"runtime/debug"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get("X-Request-ID")

		if requestID == "" {
			requestID = uuid.NewString()
		}

		c.Locals("requestID", requestID)
		c.Set("X-Request-ID", requestID)

		return c.Next()
	}
}

// identify the request and response (api) ////

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		requestID, _ := c.Locals("requestID").(string)

		log.Printf(
			"Method=%s Path=%s Status=%d Duration=%s RequestID=%s",
			c.Method(),
			c.Path(),
			c.Response().StatusCode(),
			duration,
			requestID,
		)

		return err
	}
}

// identifying the panics ////

func Recovery() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				log.Printf(
					"Panic recovered: %v\n%s",
					err,
					debug.Stack(),
				)

				_ = c.Status(fiber.StatusInternalServerError).JSON(
					fiber.Map{
						"success": false,
						"message": "Internal server error",
					},
				)
			}
		}()

		return c.Next()
	}
}
