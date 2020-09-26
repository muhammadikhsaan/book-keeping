package authorization

import (
	"Accounting/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	service = services.NewLoginService()
)

//ValidationRequest used to validation reguest header
func ValidationRequest(c *gin.Context) {
	var h HeaderModel

	c.Header("Accept", "application/json")
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Methods", "*")

	if err := c.ShouldBindHeader(&h); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !service.HandleAuthenticator(h.Authorization) {
		c.AbortWithStatus(http.StatusNetworkAuthenticationRequired)
		return
	}
}
