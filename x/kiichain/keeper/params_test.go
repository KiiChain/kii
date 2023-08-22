package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "kiichain/testutil/keeper"
	"kiichain/x/kiichain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.KiichainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
