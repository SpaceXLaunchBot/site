package api

import (
	"net/http"
)

// HasSession allows the frontend to determine if the user has an active and valid session or not.
func (a Api) HasSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	exists, _, err := a.getSessionFromCookie(r)
	if err != nil || !exists {
		endWithResponse(w, responseNoSession)
	} else {
		endWithResponse(w, responseAllOk)
	}
}
