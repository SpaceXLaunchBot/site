package api

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

// Can be anything, having it as route makes sense to me.
const statsCacheKey string = "/api/stats"

// Stats is the endpoint handler for statistics derived from collected metrics.
func (a Api) Stats(c *gin.Context) {
	if cached, ok := a.cache.Get(statsCacheKey); ok {
		if cachedResponse, ok := cached.(statsResponse); ok {
			endWithResponse(c, &cachedResponse)
			return
		}
	}

	countRecords, actionCounts, err := a.db.Stats()
	if err != nil {
		endWithResponse(c, responseDatabaseError)
		return
	}

	resp := statsResponse{}
	resp.Success = true
	resp.CountRecords = countRecords
	resp.ActionCounts = actionCounts
	a.cache.Set(statsCacheKey, resp, cache.DefaultExpiration)
	endWithResponse(c, &resp)
}
