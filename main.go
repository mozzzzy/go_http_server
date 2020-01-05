package main

/*
 * Module Dependencies
 */
import (
	"fmt"
	"os"

	"github.com/mozzzzy/arguments"
)

/*
 * Types
 */

/*
 * Constants
 */
const DEFAULT_ADDR_STR string = ""
const DEFAULT_PORT_HTTP uint = 80
const DEFAULT_PORT_HTTPS uint = 443
const DEFAULT_READ_TIMEOUT_SEC = 10
const DEFAULT_WRITE_TIMEOUT_SEC = 10
const DEFAULT_MAX_HEADER_BYTES = 1 << 20
const DEFAULT_KEY_FILE_PATH = "./key.pem"
const DEFAULT_CERT_FILE_PATH = "./cert.pem"

/*
 * Functions
 */

func getArgOps() (arguments.Arguments, error) {
	optionRules := []arguments.Option{
		{
			LongKey:     "help",
			ShortKey:    "h",
			ValueType:   "bool",
			Description: "Show usage message and exit.",
		},
	}

	var args arguments.Arguments
	args.AddRules(optionRules)
	parseErr := args.Parse()

	return args, parseErr
}

func main() {
	args, parseErr := getArgOps()
	if parseErr != nil {
		fmt.Fprintf(
			os.Stderr,
			"Failed to parse argument options. %v\n",
			parseErr,
		)
		args.Usage()
		os.Exit(1)
	}

	if args.IsTrue("h") || args.IsTrue("help") {
		args.Usage()
		os.Exit(0)
	}

	// Create httpServer
	httpServer := NewHttpServer(
		DEFAULT_ADDR_STR,
		DEFAULT_PORT_HTTP,
		DEFAULT_READ_TIMEOUT_SEC,
		DEFAULT_WRITE_TIMEOUT_SEC,
		DEFAULT_MAX_HEADER_BYTES,
	)

	// Create httpsServer
	httpsServer := NewHttpServer(
		DEFAULT_ADDR_STR,
		DEFAULT_PORT_HTTPS,
		DEFAULT_READ_TIMEOUT_SEC,
		DEFAULT_WRITE_TIMEOUT_SEC,
		DEFAULT_MAX_HEADER_BYTES,
	)

	// Add request handler
	var reqHandler requestHandler
	AddHandler(reqHandler)

	// Create goroutine and listen http port
	go httpServer.ListenAndServe()
	fmt.Printf(
		"Http server goroutine listens at %v:%v\n",
		DEFAULT_ADDR_STR,
		DEFAULT_PORT_HTTP)

	// Create goroutine and listen https port
	go httpsServer.ListenAndServeTLS(DEFAULT_CERT_FILE_PATH, DEFAULT_KEY_FILE_PATH)
	fmt.Printf(
		"Https server goroutine listens at %v:%v\n",
		DEFAULT_ADDR_STR,
		DEFAULT_PORT_HTTPS)

	// Wait for httpServer and httpsServer to finish
	select {
	case httpServerErr := <-httpServer.FinishChan:
		if httpServerErr != nil {
			fmt.Printf(
				"Http server goroutine stoped with error. %v\n",
				httpServerErr)
			os.Exit(1)
		}
	case httpsServerErr := <-httpsServer.FinishChan:
		if httpsServerErr != nil {
			fmt.Printf(
				"Https server goroutine stoped with error. %v\n",
				httpsServerErr)
			os.Exit(1)
		}
	}
	os.Exit(0)
}
