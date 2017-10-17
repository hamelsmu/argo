package cmd

import (
	"fmt"

	"github.com/argoproj/argo"
	"github.com/spf13/cobra"
)

const (
	// CLIName is the name of the CLI
	CLIName = "argo"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: fmt.Sprintf("Print the version number of %s", CLIName),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s version %s\n", CLIName, argo.FullVersion)
	},
}
