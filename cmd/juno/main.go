package main

import (
	"os"

	"github.com/forbole/juno/cmd/parse/types"

	"github.com/forbole/juno/modules/messages"
	"github.com/forbole/juno/modules/registrar"

	"github.com/forbole/juno/cmd"
)

func main() {
	// JunoConfig the runner
	config := cmd.NewConfig("juno").
		WithParseConfig(types.NewConfig().
			WithRegistrar(registrar.NewDefaultRegistrar(
				messages.CosmosMessageAddressesParser,
			)),
		)

	// Run the commands and panic on any error
	exec := cmd.BuildDefaultExecutor(config)
	err := exec.Execute()
	if err != nil {
		os.Exit(1)
	}
}
