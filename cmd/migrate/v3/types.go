package v3

import (
	databaseconfig "github.com/forbole/juno/database/config"
	loggingconfig "github.com/forbole/juno/logging/config"
	"github.com/forbole/juno/modules/pruning"
	"github.com/forbole/juno/modules/telemetry"
	nodeconfig "github.com/forbole/juno/node/config"
	parserconfig "github.com/forbole/juno/parser/config"
	pricefeedconfig "github.com/forbole/juno/pricefeed"
	"github.com/forbole/juno/types/config"
)

// Config defines all necessary juno configuration parameters.
type Config struct {
	Chain    config.ChainConfig    `yaml:"chain"`
	Node     nodeconfig.Config     `yaml:"node"`
	Parser   parserconfig.Config   `yaml:"parsing"`
	Database databaseconfig.Config `yaml:"database"`
	Logging  loggingconfig.Config  `yaml:"logging"`

	// The following are there to support modules which config are present if they are enabled

	Telemetry *telemetry.Config       `yaml:"telemetry,omitempty"`
	Pruning   *pruning.Config         `yaml:"pruning,omitempty"`
	PriceFeed *pricefeedconfig.Config `yaml:"pricefeed,omitempty"`
}
