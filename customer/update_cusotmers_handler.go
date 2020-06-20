package customer

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/pasuts/goonline/database"
	
)

func updateCustomersHandler(c *gin.Context) {
	id := c.Param("id")
	cust := &Customer{}

	if err := c.ShouldBindJSON(cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := database.Conn().Prepare("UPDATE customers SET status=$2, name=$3, email=$4 WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err := stmt.Exec(id, cust.Status, cust.Name, cust.Email); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, cust)
}