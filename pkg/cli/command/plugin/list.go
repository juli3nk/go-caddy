package plugin

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

//	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-caddy/plugins"
	"github.com/spf13/cobra"
)

var pluginListFilter []string

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List plugins",
		Long:    listDescription,
		Args:    cobra.NoArgs,
		Run:     runList,
	}

	flags := cmd.Flags()
	flags.StringSliceVarP(&pluginListFilter, "filter", "f", []string{}, "Filter output based on conditions provided")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
//	filters := utils.ConvertSliceToMap("=", pluginListFilter)

	p, err := plugins.New()
	if err != nil {
		log.Fatal(err)
	}

	result := p.ListAllPlugins()

	if len(*result) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tTYPE")

		for _, plg := range *result {
			fmt.Fprintf(w, "%s\t%s\n", plg.Name, plg.Type.Name)
		}

		w.Flush()
	}
}

var listDescription = `List plugins`
