package main

import (
	"apipattern/router"
	"net/http"
	"fmt"
	"log"

	"apipattern/config"
	"github.com/go-chi/chi"
)

func main() {
	// Read config file
	cfg := config.AppConfig()
	r := chi.NewRouter()

	router.Register(r)
	log.Printf("Start listening on localhost:8080")

	if err := http.ListenAndServe(fmt.Sprintf("%v:%v", cfg.HostName, cfg.Port), r); err != nil {
		panic(err)
	}
}
