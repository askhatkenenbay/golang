package main

import "net/http"

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
