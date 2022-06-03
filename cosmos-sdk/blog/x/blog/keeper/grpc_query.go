package keeper

import (
	"github.com/augustine/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
