package handlers

import (
	"net/http"
)

func sendError(res http.ResponseWriter, err error, code int) {
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(code)
	res.Write([]byte(err.Error()))
}
