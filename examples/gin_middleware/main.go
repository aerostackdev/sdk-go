package main

import (
	"context"
	"net/http"

	"github.com/aerostackdev/sdks/packages/go"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/operations"
	"github.com/gin-gonic/gin"
)

/*
 * Gin Middleware Example
 *
 * Demonstrates SDK usage within Gin framework handlers
 */

func main() {
	r := gin.Default()

	// Initialize SDK
	// apiKey := os.Getenv("AEROSTACK_API_KEY")
	client := aerostack.New(
	// aerostack.WithSecurity(shared.Security{APIKeyAuth: &apiKey}),
	)

	// Inject SDK into context middleware
	r.Use(func(c *gin.Context) {
		c.Set("sdk", client)
		c.Next()
	})

	r.POST("/signup", func(c *gin.Context) {
		client := c.MustGet("sdk").(*aerostack.SDK)

		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Name     string `json:"name"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := client.Authentication.AuthSignup(context.Background(), operations.AuthSignupRequestBody{
			Email:    req.Email,
			Password: req.Password,
			Name:     &req.Name,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if res.StatusCode == 201 {
			c.JSON(http.StatusCreated, res.AuthResponse)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Signup failed"})
		}
	})

	r.Run(":8080")
}
