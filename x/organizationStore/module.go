package organizationStore

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/sahith-narahari/sdk-tutorial/x/organizationStore/client/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return ModuleCdc.MustMarshalJSON(nil)
}

func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	return nil
}

func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
}

func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	//return cli.GetQueryCmd(StoreKey, cdc)
	return nil
}

func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(StoreKey, cdc)
	//return cli.GetTxCmd(StoreKey, cdc)
}


type AppModule struct {
	AppModuleBasic
	keeper Keeper
}

func NewAppModule(k Keeper) AppModule {
	return AppModule{
		AppModuleBasic{},
		k,
	}
}

func (AppModule) Name() string {
	return ModuleName
}

func (a AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (a AppModule) Route() string {
	return RouterKey
}

func (a AppModule) NewHandler() sdk.Handler {
	return NewHandler(a.keeper)
}

func (a AppModule) QuerierRoute() string {
	return ModuleName
}

func (a AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(a.keeper)
}

func (a AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

func (a AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func (a AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	return nil
}

func (a AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	return []byte{}
}
