package api

import (
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

// subscribedResponse is the API response for the stats API route.
type statsResponse struct {
	genericResponse
	CountRecords []database.CountRecord `json:"counts"`
	ActionCounts []database.ActionCount `json:"action_counts"`
}

// Stats is the endpoint handler for statistics derived from collected metrics.
func (a Api) Stats(c *gin.Context) {
	// We can cache this endpoint per IP to prevent refresh spam.
	cacheKey := c.FullPath() + c.ClientIP()
	if cachedResp, ok := a.cache.Get(cacheKey); ok {
		if cachedResponseAsserted, ok := cachedResp.(statsResponse); ok {
			endWithResponse(c, &cachedResponseAsserted)
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
	a.cache.Set(cacheKey, resp, cache.DefaultExpiration)
	endWithResponse(c, &resp)
}
