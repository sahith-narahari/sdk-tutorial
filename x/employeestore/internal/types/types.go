package types

import (
	"fmt"
	"strings"
)

type EmployeeInfo struct {
	EmployeeName string `json:"employee_name"`
}

func NewEmployeeInfo() EmployeeInfo{
	return EmployeeInfo{
	}
}

func (e EmployeeInfo) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Name: %s`, e.EmployeeName))
}