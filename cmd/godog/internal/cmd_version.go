package internal

import (
	"fmt"
	"os"

	"github.com/cucumber/godog"
	"github.com/spf13/cobra"
)

// CreateVersionCmd creates the version subcommand.
func CreateVersionCmd() cobra.Command {
	versionCmd := cobra.Command{
		Use:                   "version",
		Short:                 "Show current version",
		Run:                   versionCmdRunFunc,
		DisableFlagsInUseLine: true,
	}

	return versionCmd
}

func versionCmdRunFunc(cmd *cobra.Command, args []string) {
	fmt.Fprintln(os.Stdout, "Godog version is:", godog.Version)
	os.Exit(0)
}
