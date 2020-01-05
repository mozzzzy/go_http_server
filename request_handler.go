package main

/*
 * Module Dependencies
 */

import (
	"fmt"
	"html"
	"net/http"
)

/*
 * Types
 */

type requestHandler struct {
}

/*
 * Constants
 */

/*
 * Functions
 */

func (reqHandler requestHandler) ServeHTTP(
	writer http.ResponseWriter,
	req *http.Request) {
	fmt.Fprintf(writer, "Hello, %q", html.EscapeString(req.URL.Path))
}
