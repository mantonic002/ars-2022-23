package main

// swagger:response ResponseConfig
type ResponseConfig struct {
	// ConfigId of the config
	// in: string
	ConfigId string `json:"ConfigId"`

	// Map of entries of the config
	// in: body
	Entries map[string]string `json:"entries"`
}

// swagger:response ResponseGroup
type ResponseGroup struct {
	// GroupId of the group
	// in: string
	GroupId string `json:"GroupId"`

	// List of configs of the group
	// in: []Config
	Configs []Config `json:"configs"`
}

// swagger:response ErrorResponse
type ErrorResponse struct {
	// Error status code
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
}

// swagger:response NoContentResponse
type NoContentResponse struct{}
