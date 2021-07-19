package api

import (
	"net/http"
)

// VerifySession allows the frontend to determine if the user has an active and valid session or not.
func (a Api) VerifySession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// All we have to do is return all ok, as the session middleware will return if there is no session.
	endWithResponse(w, responseAllOk)
}
