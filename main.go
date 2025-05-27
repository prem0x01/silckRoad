package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const LISTEN_ON = ":8032"

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1>Hey Prem its started...</h1>`, nil)
	return
}

func main() {
	r := chi.NewRouter()

	r.Get("/", handlerIndex)
	log.Printf("Listening on %s", LISTEN_ON)
	http.ListenAndServe(LISTEN_ON, r)
}
