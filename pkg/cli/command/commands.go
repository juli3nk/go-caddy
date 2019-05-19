package command

import (
	"github.com/juliengk/go-caddy/pkg/cli/command/build"
	"github.com/juliengk/go-caddy/pkg/cli/command/plugin"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		build.NewCommand(),
		plugin.NewCommand(),
	)
}
