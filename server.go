package main

import (
	"ars-2022-23/ConfigStore"
	"errors"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
)

type configServer struct {
	store *ConfigStore.ConfigStore

	groupData map[string]*ConfigStore.Group // izigrava bazu podataka
}

// swagger:route POST /config/ config createConfig
// Add new config
//
// responses:
//
//	415: ErrorResponse
//	400: ErrorResponse
//	201: ResponseConfig
func (cs *configServer) createConfigHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	config, err := cs.store.Config(rt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, config)
}

// swagger:route GET /configs/ config getConfigs
// Get all configs
//
// responses:
//
//	200: []ResponseConfig
func (cs *configServer) getAllHandler(w http.ResponseWriter, req *http.Request) {
	allTasks, err := cs.store.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, allTasks)
}

// swagger:route GET /config/{ConfigId}/ config getConfigById
// Get config by config-id
//
// responses:
//
//	404: ErrorResponse
//	200: ResponseConfig
func (cs *configServer) getConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	task, err := cs.store.Get(id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, task)
}

// swagger:route DELETE /config/{ConfigId}/ config deleteConfig
// Delete config
//
// responses:
//
//	404: ErrorResponse
//	204: NoContentResponse
func (cs *configServer) delConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]

	msg, err := cs.store.Delete(id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, msg)
}

// swagger:route POST /group/ group createGroup
// Add new group
//
// responses:
//
//	415: ErrorResponse
//	400: ErrorResponse
//	201: ResponseGroup
func (cs *configServer) createGroupHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	group, err := decodeGroup(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := createId()
	group.GroupId = id
	cs.groupData[id] = group
	renderJSON(w, group)
}

// swagger:route PUT /group/{GroupId}/config/{ConfigId}/ group addConfigToGroup
// Add config to group
//
// responses:
//
//	415: ErrorResponse
//	400: ErrorResponse
//	201: ResponseGroup
func (cs *configServer) addConfigToGroup(w http.ResponseWriter, req *http.Request) {
	//groupId := mux.Vars(req)["groupId"]
	//id := mux.Vars(req)["id"]
	//task, ok := cs.groupData[id]
	//group, ook := cs.groupData[groupId]
	//if !ok || !ook {
	//	err := errors.New("key not found")
	//	http.Error(w, err.Error(), http.StatusNotFound)
	//	return
	//}
	//
	//group.Configs = append(group.Configs, *task)
	//cs.groupData[groupId] = group
	//
	//return
}

// swagger:route GET /groups/ group getGroups
// Get all groups
//
// responses:
//
//	200: []ResponseGroup
func (cs *configServer) getAllGroupsHandler(w http.ResponseWriter, req *http.Request) {
	allGroups := []*ConfigStore.Group{}
	for _, v := range cs.groupData {
		allGroups = append(allGroups, v)
	}

	renderJSON(w, allGroups)
}

// swagger:route GET /group/{GroupId}/ group getGroupById
// Get group by group-id
//
// responses:
//
//	404: ErrorResponse
//	200: ResponseGroup
func (cs *configServer) getGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, ok := cs.groupData[id]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, task)
}

// swagger:route DELETE /group/{GroupId}/ group deleteGroup
// Delete group
//
// responses:
//
//	404: ErrorResponse
//	204: NoContentResponse
func (cs *configServer) delGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	_, ok := cs.groupData[id]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	delete(cs.groupData, id)
}

// swagger:route DELETE /group/{GroupId}/config/{ConfigId}/ group deleteConfigFromGroup
// Delete config from group
//
// responses:
//
//	404: ErrorResponse
//	204: NoContentResponse
func (cs *configServer) delConfigFromGroupHandler(w http.ResponseWriter, req *http.Request) {
	groupId := mux.Vars(req)["groupId"]
	id := mux.Vars(req)["id"]
	group, ok := cs.groupData[groupId]
	if !ok {
		err := errors.New("group not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	for i, config := range group.Configs {
		if config.ConfigId == id {
			group.Configs = append(group.Configs[:i], group.Configs[i+1:]...)
			cs.groupData[groupId] = group
			return
		}
	}

	err := errors.New("config not found in group")
	http.Error(w, err.Error(), http.StatusNotFound)
	return
}

func (cs *configServer) swaggerHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./swagger.yaml")
}
