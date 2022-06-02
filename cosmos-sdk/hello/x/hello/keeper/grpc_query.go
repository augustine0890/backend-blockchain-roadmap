package keeper

import (
	"github.com/augustine/hello/x/hello/types"
)

var _ types.QueryServer = Keeper{}
