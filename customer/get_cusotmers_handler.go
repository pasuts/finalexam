package customer

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/pasuts/goonline/database"
)


func getCustomersHandler(c *gin.Context) {

	stmt, err := database.Conn().Prepare("SELECT id, name, email, status FROM customers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	customers := []Customer{}
	for rows.Next() {
		cust := Customer{}

		err := rows.Scan(&cust.ID, &cust.Name, &cust.Email, &cust.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		customers = append(customers, cust)
	}

	c.JSON(http.StatusOK, customers)
}