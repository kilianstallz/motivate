package keeper_test

import (
	"testing"

	testkeeper "github.com/kilianstallz/motivate/testutil/keeper"
	"github.com/kilianstallz/motivate/x/motivate/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MotivateKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
