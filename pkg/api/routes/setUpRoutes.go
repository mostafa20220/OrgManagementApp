package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes and sets up API routes.
func SetupRoutes(router *gin.Engine, dbClient *mongodb.MongoDBClient) {
	router.Use(middlewares.AuthMiddleware())

	router.POST("/signup", handlers.SignupHandler)
	router.POST("/signin", handlers.SigninHandler)
	router.POST("/refresh-token", handlers.RefreshTokenHandler)

	orgGroup := router.Group("/organization")
	orgGroup.Use(middlewares.JWTAuthMiddleware(dbClient)) // Middleware for organization routes

	orgGroup.POST("", handlers.CreateOrganizationHandler)
	orgGroup.GET("/:organization_id", handlers.ReadOrganizationHandler)
	orgGroup.GET("", handlers.ReadAllOrganizationsHandler)
	orgGroup.PUT("/:organization_id", handlers.UpdateOrganizationHandler)
	orgGroup.DELETE("/:organization_id", handlers.DeleteOrganizationHandler)

	orgGroup.POST("/:organization_id/invite", handlers.InviteUserToOrganizationHandler)
}
