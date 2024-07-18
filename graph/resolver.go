package graph

//go:generate go run ../generate/generate.go
import (
	_ "github.com/akafazov/gqlgen/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}
