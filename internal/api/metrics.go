package api

import "net/http"

func (a Api) Metrics(w http.ResponseWriter, r *http.Request) {
	endWithResponse(w, genericResponse{Error: "endpoint not implemented", StatusCode: http.StatusNotImplemented})
}
