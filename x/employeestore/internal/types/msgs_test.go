package types

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddEmployee(t *testing.T) {
	name := "lp"
	id := "12"

	emp := NewMsgStoreEmployee(name, id)

	require.Equal(t, emp.Route(), RouterKey)
	require.Equal(t, emp.Type(), "store_employee")
}

func TestNewMsgStoreEmployeeValidation(t *testing.T) {

	cases := []struct {
		valid bool
		tx    MsgEmployee
	}{
		{true, NewMsgStoreEmployee("lp", "12")},
		{false, NewMsgStoreEmployee("", "12")},
		{false, NewMsgStoreEmployee("lp", "")},
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

func TestGetEmployee(t *testing.T) {
	var msg = NewMsgStoreEmployee("lp", "12")
	res := msg.GetSignBytes()
	expexted := ``
	require.Equal(t, expexted, string(res))
}
