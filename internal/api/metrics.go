package api

import (
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/patrickmn/go-cache"
	"net/http"
)

// subscribedResponse is the API response for the metrics API route.
type metricsResponse struct {
	genericResponse
	CountRecords []database.CountRecord `json:"counts"`
}

// Metrics returns metric information.
func (a Api) Metrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// We can cache this endpoint per IP to prevent refresh spam.
	cacheKey := r.URL.String() + r.RemoteAddr
	if cachedResp, ok := a.cache.Get(cacheKey); ok {
		endWithResponse(w, cachedResp)
		return
	}

	metrics, err := a.db.Metrics()
	if err != nil {
		endWithResponse(w, responseDatabaseError)
		return
	}

	response := metricsResponse{}
	response.Success = true
	response.CountRecords = metrics
	a.cache.Set(cacheKey, response, cache.DefaultExpiration)
	endWithResponse(w, response)
}
