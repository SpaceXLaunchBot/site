package api

import (
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/patrickmn/go-cache"
	"net/http"
)

// subscribedResponse is the API response for the stats API route.
type statsResponse struct {
	genericResponse
	CountRecords []database.CountRecord `json:"counts"`
	ActionCounts []database.ActionCount `json:"action_counts"`
}

// Stats is the endpoint handler for statistics derived from collected metrics.
func (a Api) Stats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// We can cache this endpoint per IP to prevent refresh spam.
	cacheKey := r.URL.String() + r.RemoteAddr
	if cachedResp, ok := a.cache.Get(cacheKey); ok {
		if cachedResponseAsserted, ok := cachedResp.(statsResponse); ok {
			endWithResponse(w, &cachedResponseAsserted)
			return
		}
	}

	countRecords, actionCounts, err := a.db.Stats()
	if err != nil {
		endWithResponse(w, responseDatabaseError)
		return
	}

	response := statsResponse{}
	response.Success = true
	response.CountRecords = countRecords
	response.ActionCounts = actionCounts
	a.cache.Set(cacheKey, response, cache.DefaultExpiration)
	endWithResponse(w, &response)
}
