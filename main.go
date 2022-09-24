package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Peyton232/todo/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// change to env vars maybe
var port int = 8080
var host string = "localhost"

func main() {
	var port = flag.Int("port", port, "Port for test HTTP server")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	handler, err := api.NewHandler()
	if err != nil {
		panic(err)
	}

	// This is how you set up a basic chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	//r.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	api.HandlerFromMux(handler, r)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
