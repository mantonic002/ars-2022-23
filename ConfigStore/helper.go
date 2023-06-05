package ConfigStore

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	configs       = "configs/%s/%s"
	configsLabels = "configs/%s/%s/%s"
	all           = "configs"

	groups    = "groups/%s/%s"
	allGroups = "groups"
)

func generateKey(version string, labels string) (string, string) {
	id := uuid.New().String()
	if labels != "" {
		return fmt.Sprintf(configsLabels, id, version, labels), id
	} else {
		return fmt.Sprintf(configs, id, version), id
	}

}

func constructKey(id string, version string, labels string) string {
	if labels != "" {
		return fmt.Sprintf(configsLabels, id, version, labels)
	} else {
		return fmt.Sprintf(configs, id, version)
	}

}

func generateKeyGroup(version string) (string, string) {
	id := uuid.New().String()
	return fmt.Sprintf(groups, id, version), id

}

func constructKeyGroup(id string, version string) string {
	return fmt.Sprintf(groups, id, version)
}
