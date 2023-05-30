package main

import (
	"ars-2022-23/ConfigStore"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
)

func decodeBody(r io.Reader) (*ConfigStore.Config, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var c ConfigStore.Config
	if err := dec.Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func decodeGroup(r io.Reader) (*ConfigStore.Group, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var g ConfigStore.Group
	if err := dec.Decode(&g); err != nil {
		return nil, err
	}
	return &g, nil
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func createId() string {
	return uuid.New().String()
}
