package keeper

import (
	"github.com/starify-labs/starify/x/starify/types"
)

var _ types.QueryServer = Keeper{}
