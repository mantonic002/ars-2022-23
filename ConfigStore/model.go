package ConfigStore

// swagger:model Config
type Config struct {
	// ConfigId of the config
	// in: string
	ConfigId string `json:"ConfigId"`

	// Labels of the config
	// in: string
	Labels string `json:"labels"`

	// Version of the config
	// in: string
	Version string `json:"version"`

	// Map of entries of the config
	// in: map[string]string
	Entries map[string]string `json:"entries"`
}

// swagger:model Group
type Group struct {
	// GroupId of the group
	// in: string
	GroupId string `json:"GroupId"`

	// Version of the group
	// in: string
	Version string `json:"version"`

	// List of configs of the group
	// in: []Config
	Configs []Config `json:"configs"`
}
