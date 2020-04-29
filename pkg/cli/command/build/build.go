package build

import (
	"log"

	"github.com/juli3nk/go-caddy/build"
	"github.com/juli3nk/go-caddy/pkg/config"
	"github.com/juli3nk/go-caddy/plugins"
	"github.com/spf13/cobra"
)

const CONFIGFILE = ".caddybuild.yml"

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build Caddy binary",
		Long:  buildDescription,
		Args:  cobra.NoArgs,
		Run:   runBuild,
	}

	return cmd
}

func runBuild(cmd *cobra.Command, args []string) {
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

var buildDescription = `
Build Caddy binary

`
