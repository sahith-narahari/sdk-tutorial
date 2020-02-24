package types

type Org struct {
	Name string         `json:"name"`
	Address string `json:"address"`
	Employees []Employee `json:"employees"`
}

type Employee struct{
	Name string `json:"name"`
	Address string `json:"address"`
}