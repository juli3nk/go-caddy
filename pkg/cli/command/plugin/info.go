package plugin

import (
	"fmt"
	"log"

	"github.com/juliengk/go-caddy/plugins"
	"github.com/spf13/cobra"
)

func newInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info [name]",
		Short: "Get plugin information",
		Long:  infoDescription,
		Args:  cobra.ExactArgs(1),
		Run:   runInfo,
	}

	return cmd
}

func runInfo(cmd *cobra.Command, args []string) {
	p, err := plugins.New()
	if err != nil {
		log.Fatal(err)
	}

	result, err := p.GetPlugin(args[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:", result.Name)
	fmt.Println("Type:", result.Type.Name)
	fmt.Println("Description:", result.Description)
	fmt.Println("Website:", result.Website)
	fmt.Println("Docs:", result.Docs)
}

var infoDescription = `
Get plugin information

`
