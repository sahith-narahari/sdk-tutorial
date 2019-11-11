package keeper

import (
	"testing"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

var name = "maTurtle"

func TestNewMsgSet(t *testing.T)  {
	id := "1"

	coinKeeper := bank.Keeper()
	var msg = NewKeeper()
	res := msg.GetSigners()
}
