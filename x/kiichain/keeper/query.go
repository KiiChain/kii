package keeper

import (
	"kiichain/x/kiichain/types"
)

var _ types.QueryServer = Keeper{}
