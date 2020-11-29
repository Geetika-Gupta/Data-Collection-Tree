package model

//WebResponse : web response model
type WebResponse struct {
	Dimension
	Metrics
}

//DimensionResp : dimension response model
type DimensionResp struct {
	Dim []Data `json:"dim"`
}

//MetricsResp : metrics response model
type MetricsResp struct {
	Met []Data `json:"metrics"`
}

