package v3

import (
	databaseconfig "github.com/Fueav/juno/database/config"
	loggingconfig "github.com/Fueav/juno/logging/config"
	"github.com/Fueav/juno/modules/pruning"
	"github.com/Fueav/juno/modules/telemetry"
	nodeconfig "github.com/Fueav/juno/node/config"
	parserconfig "github.com/Fueav/juno/parser/config"
	pricefeedconfig "github.com/Fueav/juno/pricefeed"
	"github.com/Fueav/juno/types/config"
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
