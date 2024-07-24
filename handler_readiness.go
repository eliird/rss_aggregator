package main

import "net/http"

func handler_readiness(w http.ResponseWriter, r *http.Request) {
	type readiness struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, 200, readiness{Status: "ok"})
}
