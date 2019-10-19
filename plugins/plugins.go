package plugins

import (
	"fmt"

	"github.com/juliengk/go-caddy/pkg/utils"
	"github.com/juliengk/go-utils/json"
)

const URL = "https://caddyserver.com/v1/api/download-page"

func New() (*API, error) {
	data, err := utils.GetURL(URL)
	if err != nil {
		return nil, err
	}

	api := new(API)

	if err := json.Decode(data, api); err != nil {
		return nil, err
	}

	return api, nil
}

func (a *API) ListAllPlugins() *Plugins {
	return &a.Plugins
}

func (a *API) ListPlugins(pType string) (*Plugins, error) {
	var result Plugins

	for _, plugin := range a.Plugins {
		if plugin.Type.ID == pType {
			result = append(result, plugin)
		}
	}

	return &result, nil
}

func (a *API) GetPlugin(name string) (*Plugin, error) {
	for _, plugin := range a.Plugins {
		if plugin.Name == name {
			return &plugin, nil
		}
	}

	return nil, fmt.Errorf("Plugin %s not found", name)
}
