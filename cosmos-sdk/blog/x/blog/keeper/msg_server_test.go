package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/augustine/blog/testutil/keeper"
	"github.com/augustine/blog/x/blog/keeper"
	"github.com/augustine/blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlogKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
