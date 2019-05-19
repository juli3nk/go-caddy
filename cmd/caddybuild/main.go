package main

import (
	"github.com/juliengk/go-caddy/pkg/cli/command"
)

func main() {
	cmd := command.NewCommand()
	cmd.Execute()
}
