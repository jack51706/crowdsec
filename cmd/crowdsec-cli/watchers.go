package main

import (
	"github.com/crowdsecurity/crowdsec/pkg/apiclient"

	"github.com/spf13/cobra"
)

var listAll bool

func NewListCmd() *cobra.Command {
	/* ---- LIST COMMAND */
	var cmdList = &cobra.Command{
		Use:   "watchers",
		Short: "Watchers management",
		Long: `
Manage watchers.
You can list, add or delete some watchers.
List enabled configurations (parser/scenarios/collections) on your host.

It is possible to list also configuration from [Crowdsec Hub](https://hub.crowdsec.net) with the '-a' options.

[type] must be parsers, scenarios, postoverflows, collections
		`,
		Example: `cscli watchers <command>`,
		Args:    cobra.ExactArgs(0),
	}
	cmdList.PersistentFlags().BoolVarP(&listAll, "all", "a", false, "List as well disabled items")

	var cmdListWatchers = &cobra.Command{
		Use:   "list",
		Short: "List enabled watchers",
		Long:  ``,
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			client := &apiclient.ApiClient{}
		},
	}
	cmdList.AddCommand(cmdListWatchers)

	return cmdList
}
