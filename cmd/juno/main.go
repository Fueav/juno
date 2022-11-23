package main

import (
	"os"

	"github.com/Fueav/juno/cmd/parse/types"

	"github.com/Fueav/juno/modules/messages"
	"github.com/Fueav/juno/modules/registrar"

	"github.com/Fueav/juno/cmd"
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
