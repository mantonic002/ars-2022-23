package ConfigStore

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"os"
)

type GroupStore struct {
	cli *api.Client
}

func NewGroup() (*GroupStore, error) {
	db := os.Getenv("DB")
	dbport := os.Getenv("DBPORT")

	group := api.DefaultConfig()
	group.Address = fmt.Sprintf("%s:%s", db, dbport)
	client, err := api.NewClient(group)
	if err != nil {
		return nil, err
	}

	return &GroupStore{
		cli: client,
	}, nil
}

func (ps *GroupStore) GetGroup(id string, version string) ([]*Group, error) {
	kv := ps.cli.KV()

	data, _, err := kv.List(constructKey(id, version, ""), nil)
	if err != nil {
		return nil, err
	}

	groups := []*Group{}
	for _, pair := range data {
		group := &Group{}
		err = json.Unmarshal(pair.Value, group)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (ps *GroupStore) GetAllGroups() ([]*Group, error) {
	kv := ps.cli.KV()
	data, _, err := kv.List(all, nil)
	if err != nil {
		return nil, err
	}

	var groups []*Group
	for _, pair := range data {
		group := &Group{}
		err = json.Unmarshal(pair.Value, group)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func (ps *GroupStore) DeleteGroup(id string, version string) (map[string]string, error) {
	kv := ps.cli.KV()
	_, err := kv.DeleteTree(constructKey(id, version, ""), nil)
	if err != nil {
		return nil, err
	}

	return map[string]string{"Deleted": id}, nil
}

func (ps *GroupStore) Group(group *Group) (*Group, error) {
	kv := ps.cli.KV()

	sid, rid := generateKeyGroup(group.Version)
	group.GroupId = rid

	data, err := json.Marshal(group)
	if err != nil {
		return nil, err
	}

	p := &api.KVPair{Key: sid, Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		return nil, err
	}

	return group, nil
}
