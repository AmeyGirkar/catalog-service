package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Placeholder for MySQL connection
	// dbUrl := os.Getenv("DB_URL")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service":  "Catalog Service",
			"status":   "running",
			"database": "MySQL",
		})
	})

	r.GET("/catalog", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"products": []string{"Laptop", "Mouse", "Keyboard"},
		})
	})

	r.Run(":8081")
}
