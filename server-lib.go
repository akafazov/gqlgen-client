package main

const libdefaultPort = "8080"

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = libdefaultPort
// 	}

// 	c := graph.Config{Resolvers: &graph.Resolver{}}
// 	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", srv)

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// 	log.Fatal(http.ListenAndServe(":"+port, nil))
// }
