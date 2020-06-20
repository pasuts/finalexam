package customer

import (
	"log"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	err := createTable()
	if err != nil {
		log.Fatal("no customer table", err)
	}

	r := gin.Default()

	apiCustomer := r.Group("/api/customer")

	apiCustomer.Use(authMiddleware)

	apiCustomer.GET("/customers", getCustomersHandler)
	apiCustomer.GET("/customers/:id", getCustomerByIdHandler)
	apiCustomer.POST("/customers", createCustomersHandler)
	apiCustomer.PUT("/customers/:id", updateCustomersHandler)
	apiCustomer.DELETE("/customers/:id", deleteCustomersHandler)

	return r
}
