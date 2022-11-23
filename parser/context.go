package parser

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"

	"github.com/forbole/juno/logging"
	"github.com/forbole/juno/node"

	"github.com/forbole/juno/database"
	"github.com/forbole/juno/modules"
)

// Context represents the context that is shared among different workers
type Context struct {
	EncodingConfig *params.EncodingConfig
	Node           node.Node
	Database       database.Database
	Logger         logging.Logger
	Modules        []modules.Module
}

// NewContext builds a new Context instance
func NewContext(
	encodingConfig *params.EncodingConfig, proxy node.Node, db database.Database,
	logger logging.Logger, modules []modules.Module,
) *Context {
	return &Context{
		EncodingConfig: encodingConfig,
		Node:           proxy,
		Database:       db,
		Modules:        modules,
		Logger:         logger,
	}
}
