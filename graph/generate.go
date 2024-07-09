// go:build ignore

package graph

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	// cfg, err := config.LoadConfig("gqlgen.yml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	// err = api.Generate(cfg,
	// 	api.AddPlugin(yourplugin.New()), // This is the magic line
	// )
	err = api.Generate(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
