package main

/*
 * Module Dependencies
 */

import (
	"fmt"
	"net/http"
	"time"
)

/*
 * Types
 */

type HttpServer struct {
	server     *http.Server
	FinishChan chan error
}

/*
 * Constants
 */

/*
 * Functions
 */

func AddHandler(handler http.Handler) {
	http.Handle("/", handler)
}

func (httpServer *HttpServer) ListenAndServe() {
	err := httpServer.server.ListenAndServe()
	httpServer.FinishChan <- err
}

func (httpsServer *HttpServer) ListenAndServeTLS(certFile string, keyFile string) {
	err := httpsServer.server.ListenAndServeTLS(certFile, keyFile)
	httpsServer.FinishChan <- err
}

func NewHttpServer(
	addrStr string,
	port uint,
	readTimeoutSec time.Duration,
	writeTimeoutSec time.Duration,
	maxHeaderBytes int,
) *HttpServer {
	// Join addrStr and port by ":"
	addrAndPortStr := fmt.Sprintf("%v:%v", addrStr, port)

	// Create pointer of HttpServer structure.
	httpServer := new(HttpServer)
	// Create pointer of inner http.Server
	httpServer.server = &http.Server{
		Addr:           addrAndPortStr,
		ReadTimeout:    readTimeoutSec * time.Second,
		WriteTimeout:   writeTimeoutSec * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}
	// Create channel
	httpServer.FinishChan = make(chan error)
	return httpServer
}
