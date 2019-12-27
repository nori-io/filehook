package main

import (
	"context"

	"github.com/nori-io/nori-common/config"
	"github.com/nori-io/nori-common/logger"
	"github.com/nori-io/nori-common/meta"
)

type plugin struct {
	instance logger.Hook
	log      logger.Logger
	config   *pluginConfig
}

type pluginConfig struct {
	Path config.String
}

type instance struct {
	config *pluginConfig
}

var (
	Plugin plugin
)

func (p *plugin) Init(_ context.Context, config config.Manager) error {
	cm := config.Register(p.Meta())
	p.config = &pluginConfig{
		Path: cm.String("hooks.filehook.path", "filepath of filehook"),
	}
	return nil
}

func (p *plugin) Instance() interface{} {
	return p.instance
}

func (p plugin) Meta() meta.Meta {
	return &meta.Data{
		ID: meta.ID{
			ID:      "nori/filehook",
			Version: "1.0.0",
		},
		Author: meta.Author{
			Name: "Nori",
			URI:  "https://nori.io/",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Dependencies: []meta.Dependency{},
		Description:  meta.Description{},
		//Interface:    interfaces.HttpInterface,
		License: []meta.License{
			{
				Title: "",
				Type:  "GPLv3",
				URI:   "https://www.gnu.org/licenses/",
			},
		},
		Tags: []string{"hook", "hooks", "logger", "nori"},
	}
}
