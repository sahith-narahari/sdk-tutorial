package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var name = "maTurtle"

func TestMsgSetName(t *testing.T) {
	value := "1"
	acc := sdk.AccAddress([]byte("me"))
	var msg = NewEmpStore(name, value, acc)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "set_name")
}

func TestNewEmpStoreValidation(t *testing.T) {
	value := "1"
	acc := sdk.AccAddress([]byte("me"))
	name2 := "a"
	value2 := "2"
	acc2 := sdk.AccAddress([]byte("you"))

	cases := []struct {
		valid bool
		tx    StoreEmp
	}{
		{true, NewEmpStore(name, value, acc)},
		{true, NewEmpStore(name2, value2, acc2)},
		{true, NewEmpStore(name2, value, acc2)},
		{true, NewEmpStore(name2, value2, acc)},
		{false, NewEmpStore(name, value2, nil)},
		{false, NewEmpStore("", value2, acc2)},
		{false, NewEmpStore(name, "", acc2)},
	}

	for _, tc := range cases {
		err := tc.tx.ValidateBasic()
		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestNewEmpStoreGetSignBytes(t *testing.T) {
	id := "1"
	acc := sdk.AccAddress([]byte("me"))

	var msg = NewEmpStore(name, id, acc)
	res := msg.GetSignBytes()

	expected := `{"type":"employeestore/SetName","value":{"id":"1","name":"maTurtle","owner":"cosmos1d4js690r9j"}}`

	require.Equal(t, expected, string(res))
}

func TestNewEmpStoreGetSigners(t *testing.T)  {
	id := "1"
	acc := sdk.AccAddress([]byte("me"))

	var msg = NewEmpStore(name, id, acc)
	res := msg.GetSigners()

	var expected []sdk.AccAddress

	expected = append(expected, acc)

	require.Equal(t, expected, res)
}