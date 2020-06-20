package customer

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/pasuts/goonline/database"
)

func deleteCustomersHandler(c *gin.Context) {
	id := c.Param("id")

	err := deleteCustomerById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	
	c.JSON(http.StatusOK, "deleted customer.")
}

func deleteCustomerById(id string) error {
	stmt, err := database.Conn().Prepare("DELETE FROM customers WHERE id = $1")
	if err != nil {
		return fmt.Errorf("can't prepare delete statement: %w", err)
	}

	if _, err := stmt.Exec(id); err != nil {
		return fmt.Errorf("can't execute delete statment: %w", err)
	}

	return nil
}