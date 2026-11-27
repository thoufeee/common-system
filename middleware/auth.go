package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/thoufeee/common-system/response"
)

const (
	UserIDKey   contextKey = "user_id"
	UserRoleKey contextKey = "user_role"
)

// JWT claims
type AccessClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`

	jwt.RegisteredClaims
}

// Authentication validates the JWT access token

func Authentication(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		var tokenString string

		//   check the Authorization header
		authHeader := c.GetHeader("Authorization")

		if authHeader != "" {
			parts := strings.SplitN(authHeader, " ", 2)

			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				unauthorized(c, "Invalid authorization format")
				return
			}

			tokenString = strings.TrimSpace(parts[1])
		} else {
			// If no Authorization header, check the access cookie
			cookieToken, err := c.Cookie("access_token")
			if err != nil {
				unauthorized(c, "Access token is required")
				return
			}

			tokenString = cookieToken
		}

		if tokenString == "" {
			unauthorized(c, "Access token is required")
			return
		}

		//  Validate JWT
		claims := &AccessClaims{}

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (any, error) {
				return []byte(jwtSecret), nil
			},
			jwt.WithValidMethods([]string{
				jwt.SigningMethodHS256.Alg(),
			}),
		)

		if err != nil || token == nil || !token.Valid {
			unauthorized(c, "Invalid or expired access token")
			return
		}

		// Validate required claim
		userID := claims.UserID

		if userID == "" {
			userID = claims.Subject
		}

		if userID == "" || claims.Role == "" || claims.ExpiresAt == nil {
			unauthorized(c, "Invalid access token claims")
			return
		}

		// Store authenticated identity in Gin context
		c.Set(string(UserIDKey), userID)
		c.Set(string(UserRoleKey), claims.Role)

		ctx := context.WithValue(
			c.Request.Context(),
			UserIDKey,
			userID,
		)

		ctx = context.WithValue(
			ctx,
			UserRoleKey,
			claims.Role,
		)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// RequireRoles checks whether the authenticated user
// has one of the permitted roles
func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Authentication middleware must run first
		roleValue, exists := c.Get(string(UserRoleKey))

		if !exists {
			c.AbortWithStatusJSON(
				response.StatusUnauthorized,
				response.Error(
					response.CodeUnauthorized,
					"Authentication required",
				),
			)
			return
		}

		userRole, ok := roleValue.(string)

		if !ok || userRole == "" {
			c.AbortWithStatusJSON(
				response.StatusUnauthorized,
				response.Error(
					response.CodeForbidden,
					"Invalid user identity",
				),
			)
			return
		}

		// Check permitted roles.
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				c.Next()
				return
			}
		}

		// User does not have permission.
		c.AbortWithStatusJSON(
			response.StatusForbidden,
			response.Error(
				response.CodeForbidden,
				"You don't have permission to access this resource",
			),
		)
	}
}

// Return unauthorized response
func unauthorized(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		response.StatusUnauthorized,
		response.Error(
			response.CodeUnauthorized,
			message,
		),
	)
}

// Get authenticated user ID from context
func GetUserID(ctx context.Context) string {
	userID, _ := ctx.Value(UserIDKey).(string)
	return userID
}

// Get authenticated user role from context.
func GetUserRole(ctx context.Context) string {
	role, _ := ctx.Value(UserRoleKey).(string)
	return role
}
