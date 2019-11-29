package gopfserver

import (
	"github.com/gin-gonic/gin"
	"github.com/kanhaiya15/gopf/controllers"
	"github.com/kanhaiya15/gopf/middlewares"
)

// NewRouter API Route
func NewRouter() *gin.Engine {
	router := gin.New()
	router.NoRoute(func(c *gin.Context) {
		controllers.NotFound(c)
	})
	router.Use(CORSMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	route := router.Group("/gopf/api/v1")

	route.GET("/", controllers.Home)
	route.GET("/ping", controllers.Ping)
	route.GET("/health", controllers.Health)
	route.GET("/stats", controllers.DBStats)

	// jsonplaceholder
	route.GET("/posts", controllers.GetPosts)
	route.POST("/post", controllers.AddPost)
	route.PUT("/post/:postID", controllers.UpdatePost)
	route.DELETE("/post/:postID", controllers.DeletePost)

	route.Use(middlewares.Authenticate())
	return router
}

// CORSMiddleware CORS Middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method != "OPTIONS" {
			c.Next()
		}
		c.AbortWithStatus(204)
	}
}
