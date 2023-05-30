package main

import "ars-2022-23/ConfigStore"

// swagger:parameters deleteConfig
type DeleteConfigRequest struct {
	// Config ConfigId
	// in: path
	ConfigId string `json:"ConfigId"`
}

// swagger:parameters getConfigById
type GetConfigRequest struct {
	// Config ConfigId
	// in: path
	ConfigId string `json:"ConfigId"`
}

// swagger:parameters config createConfig
type RequestConfigBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	Body ConfigStore.Config `json:"body"`
}

// swagger:parameters config createGroup
type RequestGroupBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	Body ConfigStore.Group `json:"body"`
}

// swagger:parameters addConfigToGroup
type AddConfigToGroupRequest struct {
	// Group GroupId
	// in: path
	GroupId string `json:"GroupId"`
	// Config ConfigId
	// in: path
	ConfigId string `json:"ConfigId"`
}

// swagger:parameters getGroupById
type GetGroupRequest struct {
	// Group GroupId
	// in: path
	GroupId string `json:"GroupId"`
}

// swagger:parameters deleteGroup
type DeleteGroupRequest struct {
	// Group GroupId
	// in: path
	GroupId string `json:"GroupId"`
}

// swagger:parameters deleteConfigFromGroup
type DeleteConfigFromGroupRequest struct {
	// Group GroupId
	// in: path
	GroupId string `json:"GroupId"`
	// in: path
	// Config ConfigId
	ConfigId string `json:"ConfigId"`
}
