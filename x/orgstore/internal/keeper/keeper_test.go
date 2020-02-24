package keeper

import (
	"github.com/cosmos/cosmos-sdk/x/bank"
	"nameservice/x/orgstore/internal/types"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type testInput struct {
	cdc *codec.Codec
	ctx sdk.Context
	ak  auth.AccountKeeper
	pk  params.Keeper
	ns Keeper
	bk bank.Keeper
}

func setupTestInput() testInput {
	db := dbm.NewMemDB()

	cdc := codec.New()
	auth.RegisterCodec(cdc)

	authCapKey := sdk.NewKVStoreKey(auth.StoreKey)
	fckCapKey := sdk.NewKVStoreKey(types.StoreKey)
	keyParams := sdk.NewKVStoreKey("params")
	tkeyParams := sdk.NewTransientStoreKey("transient_params")

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(authCapKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(fckCapKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, db)
	ms.LoadLatestVersion()



	pk := params.NewKeeper(cdc, keyParams, tkeyParams,params.DefaultCodespace)
	ak := auth.NewAccountKeeper(
		cdc, authCapKey, pk.Subspace(auth.DefaultParamspace),auth.ProtoBaseAccount,
	)

	m := make(map[string]bool)

	bankKeeper := bank.NewBaseKeeper(ak,pk.Subspace(auth.DefaultParamspace), params.DefaultCodespace,m)

	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	bankKeeper.SetSendEnabled(ctx, true)

	ak.SetParams(ctx, auth.DefaultParams())

	return testInput{cdc: cdc, ctx: ctx, ak: ak, pk: pk , bk:bankKeeper}
}

func TestAddOrgKeeper(t *testing.T) {


	input := setupTestInput()

	t.Log("")
	ctx := input.ctx
	addr := sdk.AccAddress([]byte("addr1"))


	acc := input.ak.NewAccountWithAddress(ctx, addr)

	// Test GetCoins/SetCoins
	input.ak.SetAccount(ctx, acc)
	require.True(t, input.bk.GetCoins(ctx, addr).IsEqual(sdk.NewCoins()))

	input.bk.SetCoins(ctx, addr, sdk.NewCoins(sdk.NewInt64Coin("stake", 100000000)))
	require.True(t, input.bk.GetCoins(ctx, addr).IsEqual(sdk.NewCoins(sdk.NewInt64Coin("stake", 100000000))))

}

