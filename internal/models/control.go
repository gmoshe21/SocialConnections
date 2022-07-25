package models

type Connections struct {
	Sender    string `json:"sender" validate:"required"`
	Recipient string `json:"recipient" validate:"required"`
}

type FetchConnections struct {
	MaxConnections     int `json:"MaxConnections" validate:"required"`
	MinConnections     int `json:"MinConnections" validate:"required"`
	AverageConnections int `json:"AverageConnections" validate:"required"`
	Connections        Connections `json:"json_agg" validate:"required"`
}