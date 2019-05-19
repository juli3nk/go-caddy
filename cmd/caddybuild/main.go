package main

import (
	"log"

	"github.com/juliengk/go-caddy/build"
	"github.com/juliengk/go-caddy/pkg/config"
	"github.com/juliengk/go-caddy/plugins"
)

const CONFIGFILE = ".caddybuild.yml"

func main() {
	var importPaths []string

	c, err := config.New(CONFIGFILE)
	if err != nil {
		log.Fatal(err)
	}

	p, err := plugins.New()
	if err != nil {
		log.Fatal(err)
	}

	for _, plugin := range c.Plugins {
		info, err := p.GetPlugin(plugin)
		if err != nil {
			log.Println(err)
		}

		importPaths = append(importPaths, info.ImportPath)
	}

	b := build.New(importPaths, false)

	if err := b.GenerateGoFile(); err != nil {
		log.Fatal(err)
	}

	if err := b.WriteGoFile(); err != nil {
		log.Fatal(err)
	}

	if err := b.Run(); err != nil {
		log.Fatal(err)
	}
}
