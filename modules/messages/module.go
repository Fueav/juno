package messages

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Fueav/juno/database"
	"github.com/Fueav/juno/modules"
	"github.com/Fueav/juno/types"
)

var _ modules.Module = &Module{}

// Module represents the module allowing to store messages properly inside a dedicated table
type Module struct {
	parser MessageAddressesParser

	cdc codec.Codec
	db  database.Database
}

func NewModule(parser MessageAddressesParser, cdc codec.Codec, db database.Database) *Module {
	return &Module{
		parser: parser,
		cdc:    cdc,
		db:     db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "messages"
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *types.Tx) error {
	return HandleMsg(index, msg, tx, m.parser, m.cdc, m.db)
}
