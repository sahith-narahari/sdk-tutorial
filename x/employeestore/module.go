package employeestore

import (
	"github.com/comdex-blockchain/x/bank"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gogo/protobuf/codec"

	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	"github.com/sahith-narahari/sdk-tutorial/x/employeestore/client/cli"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

// Get the root tx command of this module
func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(StoreKey, cdc)
}

type AppModule struct {
	AppModuleBasic
	keeper     Keeper
	coinKeeper bank.Keeper
}

func NewAppModule(k Keeper, bankKeeper bank.Keeper) AppModule {

	return AppModule{
		AppModuleBasic{},
		k,
		bankKeeper,
	}
}

func (AppModule) Name() string {
	return ModuleName
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (am AppModule) Route() string {
	return RouterKey
}

func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}

func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

func (AppModuleBasic) DefaultGenesis() json.RawMessage {

	return nil
}

func (am AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.Validator {
	return nil
}

func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.Validator {
	//var genesisState GenesisState
	//ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	return nil
}

func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	//gs := ExportGenesis(ctx, am.keeper)
	return []byte{}
}

// Get the root query command of this module
func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return nil
}

func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(am.keeper)
}

func (am AppModule) QuerierRoute() string {
	return ModuleName
}

// Register rest routes
func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	//rest.RegisterRoutes(ctx, rtr, StoreKey)
}

// Validation check of the Genesis
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {

	return nil
}
