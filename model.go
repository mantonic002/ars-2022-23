package main

// swagger:model Config
type Config struct {
	// Id of the config
	// in: string
	Id string `json:"id"`

	// Map of entries of the config
	// in: map[string]string
	Entries map[string]string `json:"entries"`
}

// swagger:model Group
type Group struct {
	// Id of the group
	// in: string
	Id string `json:"id"`

	// List of configs of the group
	// in: []Config
	Configs []Config `json:"configs"`
}
