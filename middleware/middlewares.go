package logger

import (
	"context"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type contextKey string

const loggerContextKey contextKey = "logger"

// GetLogger retrieves request-specific logger.
func GetLogger(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerContextKey).(*slog.Logger); ok {
		return logger
	}

	return slog.Default()
}

// RequestID middleware
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")

		if requestID == "" {
			requestID = uuid.NewString()
		}

		c.Set("requestID", requestID)
		c.Header("X-Request-ID", requestID)

		logger := slog.Default().With(
			"request_id", requestID,
		)

		ctx := context.WithValue(
			c.Request.Context(),
			loggerContextKey,
			logger,
		)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// Logger middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)

		logger := GetLogger(c.Request.Context())

		logger.Info(
			"HTTP request completed",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", duration.String(),
			"client_ip", c.ClientIP(),
		)
	}
}

// Recovery middleware
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		logger := GetLogger(c.Request.Context())

		logger.Error(
			"panic recovered",
			"error", recovered,
			"stack", string(debug.Stack()),
		)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"code":    "INTERNAL_SERVER_ERROR",
				"message": "Internal server error",
			},
		)
	})
}
