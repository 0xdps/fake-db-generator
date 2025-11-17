package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/0xdps/fake-stack/apimock/lib/handlers"
	"github.com/0xdps/fake-stack/apimock/lib/middleware"
	"github.com/0xdps/fake-stack/apimock/lib/schema"
	"github.com/gin-gonic/gin"
)

var (
	router   *gin.Engine
	registry *schema.Registry
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	
	registry = schema.NewRegistry()
	schemasDir := os.Getenv("SCHEMAS_DIR")
	if schemasDir == "" {
		// Try multiple paths to find schemas
		possiblePaths := []string{
			filepath.Join("shared", "schemas"),           // Vercel production (includeFiles makes it relative)
			filepath.Join("..", "shared", "schemas"),     // From api/ directory
			filepath.Join("..", "..", "..", "shared", "schemas"), // From Vercel dev cache
		}
		
		var err error
		for _, path := range possiblePaths {
			if err = registry.LoadSchemas(path); err == nil {
				break
			}
		}
		
		if err != nil {
			cwd, _ := os.Getwd()
			panic(fmt.Sprintf("Failed to load schemas. CWD: %s, Error: %v", cwd, err))
		}
	} else {
		if err := registry.LoadSchemas(schemasDir); err != nil {
			panic("Failed to load schemas: " + err.Error())
		}
	}
	
	router = gin.New()
	
	// Custom recovery middleware for 500 errors
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal server error",
			"message": "An unexpected error occurred",
			"details": fmt.Sprintf("%v", recovered),
		})
	}))
	
	router.Use(middleware.SetupCORS())
	
	// Health check endpoint
	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "apimock.codes API",
			"version":   "1.0.0",
			"docs":      "https://apimock.codes/docs",
			"resources": registry.GetAllResourceNames(),
		})
	})
	
	dynamicHandler := handlers.NewDynamicHandler(registry)
	
	resourceNames := registry.GetAllResourceNames()
	nestedResources := []string{}
	regularResources := []string{}
	
	for _, name := range resourceNames {
		routePath := registry.GetRoutePath(name)
		if strings.Contains(routePath, ":") {
			nestedResources = append(nestedResources, name)
		} else {
			regularResources = append(regularResources, name)
		}
	}
	
	// Register routes at /api prefix (Vercel passes full path)
	api := router.Group("/api")
	
	for _, resourceName := range nestedResources {
		routePath := registry.GetRoutePath(resourceName)
		api.GET(routePath, dynamicHandler.GetCollection(resourceName))
		api.GET(routePath+"/meta", dynamicHandler.GetResourceMetadata(resourceName))
		
		for _, alias := range registry.GetRouteAliases(resourceName) {
			api.GET(alias, dynamicHandler.GetCollection(resourceName))
			api.GET(alias+"/meta", dynamicHandler.GetResourceMetadata(resourceName))
		}
	}
	
	for _, resourceName := range regularResources {
		routePath := registry.GetRoutePath(resourceName)
		api.GET(routePath, dynamicHandler.GetCollection(resourceName))
		api.GET(routePath+"/:id", dynamicHandler.GetSingle(resourceName))
		api.GET(routePath+"/meta", dynamicHandler.GetResourceMetadata(resourceName))
		
		for _, alias := range registry.GetRouteAliases(resourceName) {
			api.GET(alias, dynamicHandler.GetCollection(resourceName))
			api.GET(alias+"/:id", dynamicHandler.GetSingle(resourceName))
			api.GET(alias+"/meta", dynamicHandler.GetResourceMetadata(resourceName))
		}
	}
	
	// 404 handler for unmatched routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":     "Route not found",
			"path":      c.Request.URL.Path,
			"method":    c.Request.Method,
			"message":   "The requested endpoint does not exist",
			"available": registry.GetAllResourceNames(),
		})
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
