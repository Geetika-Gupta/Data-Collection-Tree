package model

//WebRequest : web request model
type WebRequest struct {
	Dimension
	Metrics
}

//Dimension : dimension request model
type Dimension struct {
	Dim []Data `json:"dim"`
}

//Metrics : metrics request model
type Metrics struct {
	Met []Data `json:"metrics"`
}

//Data : data request model
type Data struct {
	Key string `json:"key"`
	Value string `json:"val"`
}
