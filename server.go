package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/akafazov/gqlgen-client/graph"
	libgraph "github.com/akafazov/gqlgen/graph"
)

const defaultPort = "8080"

type QueryRequest struct {
	Query string `json:"query"`
}

// isLibQuery is a custom function that determines which server to use
func isLibQuery(query string) bool {
	return !strings.Contains(query, "query todos")
}

var srv *handler.Server
var srvlib *handler.Server

func CustomHandler(w http.ResponseWriter, r *http.Request) {
	// var queryRequest QueryRequest
	// if err := json.NewDecoder(r.Body).Decode(&queryRequest); err != nil {
	// 	http.Error(w, "Bad request", http.StatusBadRequest)
	// 	return
	// }

	// get r.Body as string
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// srv.ServeHTTP(w, r)
	r.Body = io.NopCloser(bytes.NewBuffer(body)) // restore the body for further use

	if isLibQuery(string(body)) {
		srvlib.ServeHTTP(w, r)
	} else {
		srv.ServeHTTP(w, r)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv = handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srvlib = handler.NewDefaultServer(libgraph.NewExecutableSchema(libgraph.Config{Resolvers: &libgraph.Resolver{}}))
	// http.Handle("/query", srv)
	// http.Handle("/libquery", srvlib)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.HandleFunc("/query", CustomHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
