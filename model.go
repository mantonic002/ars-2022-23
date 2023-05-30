package main

// swagger:model Config
type Config struct {
	// ConfigId of the config
	// in: string
	ConfigId string `json:"ConfigId"`

	// Map of entries of the config
	// in: map[string]string
	Entries map[string]string `json:"entries"`
}

// swagger:model Group
type Group struct {
	// GroupId of the group
	// in: string
	GroupId string `json:"GroupId"`

	// List of configs of the group
	// in: []Config
	Configs []Config `json:"configs"`
}
