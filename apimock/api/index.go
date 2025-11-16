package handler

import (
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
		schemasDir = filepath.Join("..", "..", "shared", "schemas")
	}
	
	if err := registry.LoadSchemas(schemasDir); err != nil {
		schemasDir = filepath.Join("shared", "schemas")
		if err := registry.LoadSchemas(schemasDir); err != nil {
			panic("Failed to load schemas: " + err.Error())
		}
	}
	
	router = gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.SetupCORS())
	
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "apimock.codes API",
			"version":   "1.0.0",
			"docs":      "https://apimock.codes/docs",
			"resources": registry.GetAllResourceNames(),
		})
	})
	
	dynamicHandler := handlers.NewDynamicHandler(registry)
	
	api := router.Group("/api")
	{
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
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
