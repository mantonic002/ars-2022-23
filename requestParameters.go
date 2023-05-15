package main

// swagger:parameters deleteConfig
type DeleteConfigRequest struct {
	// Config ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters getConfigById
type GetConfigRequest struct {
	// Config ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters config createConfig
type RequestConfigBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/RequestPost"
	//  required: true
	Body Config `json:"body"`
}

// swagger:parameters config createGroup
type RequestGroupBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/RequestPost"
	//  required: true
	Body Group `json:"body"`
}

// swagger:parameters addConfigToGroup
type AddConfigToGroupRequest struct {
	// Group ID
	// Config ID
	// in: path
	GroupId  string `json:"group-id"`
	ConfigId string `json:"config-id"`
}

// swagger:parameters getGroupById
type GetGroupRequest struct {
	// Group ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters deleteGroup
type DeleteGroupRequest struct {
	// Group ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters deleteConfigFromGroup
type DeleteConfigFromGroupRequest struct {
	// Group ID
	// Config ID
	// in: path
	GroupId  string `json:"group-id"`
	ConfigId string `json:"config-id"`
}
