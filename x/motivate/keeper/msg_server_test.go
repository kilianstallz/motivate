package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/kilianstallz/motivate/testutil/keeper"
	"github.com/kilianstallz/motivate/x/motivate/keeper"
	"github.com/kilianstallz/motivate/x/motivate/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MotivateKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
