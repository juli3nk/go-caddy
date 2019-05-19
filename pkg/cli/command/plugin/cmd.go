package plugin

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plugin",
		Short: "Plugins information",
		Long:  description,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmd.AddCommand(
		newInfoCommand(),
		newListCommand(),
	)

	return cmd
}

var description = `The **caddybuild plugin** command has subcommands for plugins information.
To see help for a subcommand, use:
    caddybuild plugin [command] --help

`
