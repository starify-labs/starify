package keeper_test

import (
	"testing"

	testkeeper "github.com/starify-labs/starify/testutil/keeper"
	"github.com/starify-labs/starify/x/starify/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.StarifyKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
