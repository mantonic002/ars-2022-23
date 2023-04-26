package main

import (
	"errors"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
)

type configServer struct {
	data      map[string]*Config
	groupData map[string]*Group // izigrava bazu podataka
}

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

	id := createId()
	rt.Id = id
	cs.data[id] = rt
	renderJSON(w, rt)
}

func (cs *configServer) getAllHandler(w http.ResponseWriter, req *http.Request) {
	allConfigs := []*Config{}
	for _, v := range cs.data {
		allConfigs = append(allConfigs, v)
	}

	renderJSON(w, allConfigs)
}

func (cs *configServer) getConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, ok := cs.data[id]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, task)
}
func (cs *configServer) delConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	if v, ok := cs.data[id]; ok {
		delete(cs.data, id)
		renderJSON(w, v)
	} else {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

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
	group.Id = id
	cs.groupData[id] = group
	renderJSON(w, group)
}

func (cs *configServer) AddConfigToGroup(w http.ResponseWriter, req *http.Request) {
	groupId := mux.Vars(req)["groupId"]
	id := mux.Vars(req)["id"]
	task, ok := cs.data[id]
	group, ook := cs.groupData[groupId]
	if !ok || !ook {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	group.Configs = append(group.Configs, *task)
	cs.groupData[groupId] = group

	return
}

func (cs *configServer) getAllGroupsHandler(w http.ResponseWriter, req *http.Request) {
	allGroups := []*Group{}
	for _, v := range cs.groupData {
		allGroups = append(allGroups, v)
	}

	renderJSON(w, allGroups)
}

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
		if config.Id == id {
			group.Configs = append(group.Configs[:i], group.Configs[i+1:]...)
			cs.groupData[groupId] = group
			return
		}
	}

	err := errors.New("config not found in group")
	http.Error(w, err.Error(), http.StatusNotFound)
	return
}
