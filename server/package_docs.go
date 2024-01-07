// Package server contains the implementation of a server.
//
// It provides functions to start and stop the server, as well as handling requests.
// The server can listen for incoming connections on a specified port and route requests
// to the appropriate handler based on the request path.
//
// The server also supports graceful shutdown by catching interrupt signals and waiting
// for ongoing requests to finish before shutting down.
//
// The server package assumes a basic understanding of HTTP servers and handlers in Go.
package server
