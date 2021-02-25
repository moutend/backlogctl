package app

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "print version",
	RunE:    versionCommandRunE,
}

func versionCommandRunE(cmd *cobra.Command, args []string) error {
	info, ok := debug.ReadBuildInfo()

	if !ok {
		return fmt.Errorf("failed to read version")
	}

	cmd.Println(info.Main.Version)

	return nil
}

func init() {
	RootCommand.AddCommand(versionCommand)
}
