// golang-simple-server implements a simple "Hello World"-esque Golang
// HTTP server that handles requests with a server bound to a customizable port
// It's useful for testing other applications and services that need to ensure
// connectivity between servers
package main

import (
	// Standard lib
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	// Third-party
	log "github.com/Sirupsen/logrus"
)

var (
	// Port server will listen on
	port *string
)

// handle is the main HTTP handler function for all requests to the server
func handle(w http.ResponseWriter, req *http.Request) {
	// Log request handling
	log.Infof("Handling request for '%s'", req.URL.String())

	fmt.Fprintf(w, "Hi there, I'm listening on port %s!", *port)
}

// configures and starts up an HTTP server on a desired port
func startServer() {
	// Create new mux instance
	mux := http.NewServeMux()

	// Create new server instance
	server := &http.Server{}

	// Set up server
	server.Addr = ":" + *port
	server.Handler = mux
	server.ReadTimeout = time.Duration(30) * time.Second
	server.WriteTimeout = time.Duration(30) * time.Second

	// Set up server routes
	mux.Handle("/", http.HandlerFunc(handle))

	// Log server start
	log.Infof("Server running on port %s", *port)

	// Attempt to start the server
	go server.ListenAndServe()
}

func main() {
	// Log start of the service
	log.Info("Server is starting")

	// Get port from flags with fallback
	port = flag.String("port", "8080", "default server port, ex: 8080")

	// Parse flags
	flag.Parse()

	// Start server
	startServer()

	// Listen for and exit the application on SIGKILL or SIGINT
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, os.Kill)

	select {
	case <-stop:
		log.Info("Server is shutting down")
	}
}
