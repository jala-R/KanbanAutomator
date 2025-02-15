package handlers

import "net/http"

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(501)
	w.Write([]byte(err.Error()))
}
