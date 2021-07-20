package api

import (
	"github.com/gin-gonic/gin"
)

// VerifySession allows the frontend to determine if the user has an active and valid session or not.
func (a Api) VerifySession(c *gin.Context) {
	// All we have to do is return all ok, as the session middleware will return if there is no session.
	endWithResponse(c, responseAllOk)
}
