package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Routes for the API Gateway, which forwards requests to the appropriate microservices
	r.GET("/api/products", handleProductRequest)
	r.POST("/api/products", handleProductRequest)
	r.PUT("/api/products/:id", handleProductRequest)
	r.DELETE("/api/products/:id", handleProductRequest)

	r.GET("/api/orders", handleOrderRequest)
	r.POST("/api/orders", handleOrderRequest)
	r.PUT("/api/orders/:id", handleOrderRequest)
	r.DELETE("/api/orders/:id", handleOrderRequest)

	r.GET("/api/users", handleUserRequest)
	r.POST("/api/users", handleUserRequest)
	r.PUT("/api/users/:id", handleUserRequest)
	r.DELETE("/api/users/:id", handleUserRequest)

	r.GET("/api/faqs", handleFAQRequest)
	r.POST("/api/faqs", handleFAQRequest)
	r.PUT("/api/faqs/:id", handleFAQRequest)
	r.DELETE("/api/faqs/:id", handleFAQRequest)

	if err := r.Run(":9090"); err != nil {
		log.Fatal("API Gateway failed to start:", err)
	}
}

// Forward the requests to the respective service based on HTTP method
func handleProductRequest(c *gin.Context) {
	// Forward request to product service
	method := c.Request.Method
	
	productServiceURL := "http://localhost:9091/api/products"

	switch method {
	case http.MethodGet:
		c.Redirect(http.StatusFound, productServiceURL)
	case http.MethodPost:
		c.Redirect(http.StatusFound, productServiceURL)
	case http.MethodPut:
		c.Redirect(http.StatusFound, productServiceURL+"/"+c.Param("id"))
	case http.MethodDelete:
		c.Redirect(http.StatusFound, productServiceURL+"/"+c.Param("id"))
	}
}

func handleOrderRequest(c *gin.Context) {
	// Forward request to order service
	method := c.Request.Method
	orderServiceURL := "http://localhost:9092/api/orders"

	switch method {
	case http.MethodGet:
		c.Redirect(http.StatusFound, orderServiceURL)
	case http.MethodPost:
		c.Redirect(http.StatusFound, orderServiceURL)
	case http.MethodPut:
		c.Redirect(http.StatusFound, orderServiceURL+"/"+c.Param("id"))
	case http.MethodDelete:
		c.Redirect(http.StatusFound, orderServiceURL+"/"+c.Param("id"))
	}
}

func handleUserRequest(c *gin.Context) {
	// Forward request to user service
	method := c.Request.Method
	fmt.Println("Method is : ", method)
	userServiceURL := "http://localhost:9093/api/users"

	switch method {
	case http.MethodGet:
		c.Redirect(http.StatusFound, userServiceURL)
	case http.MethodPost:
		c.Redirect(http.StatusFound, userServiceURL)
	case http.MethodPut:
		c.Redirect(http.StatusFound, userServiceURL+"/"+c.Param("id"))
	case http.MethodDelete:
		c.Redirect(http.StatusFound, userServiceURL+"/"+c.Param("id"))
	}
}

func handleFAQRequest(c *gin.Context) {
	// Forward request to FAQ service
	method := c.Request.Method
	faqServiceURL := "http://localhost:9094/api/faqs"

	switch method {
	case http.MethodGet:
		c.Redirect(http.StatusFound, faqServiceURL)
	case http.MethodPost:
		c.Redirect(http.StatusFound, faqServiceURL)
	case http.MethodPut:
		c.Redirect(http.StatusFound, faqServiceURL+"/"+c.Param("id"))
	case http.MethodDelete:
		c.Redirect(http.StatusFound, faqServiceURL+"/"+c.Param("id"))
	}
}
