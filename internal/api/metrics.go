package api

import "net/http"

func (a Api) Metrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	endWithResponse(w, genericResponse{Error: "endpoint not implemented", StatusCode: http.StatusNotImplemented})
}
