package app

import (
	"github.com/moutend/backlogctl/internal/cache"
	"github.com/spf13/cobra"
)

var projectCommand = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	RunE:    projectCommandRunE,
}

func projectCommandRunE(cmd *cobra.Command, args []string) error {
	if err := cache.FetchProjects(cmd.Context()); err != nil {
		return err
	}

	return nil
}

func init() {
	RootCommand.AddCommand(projectCommand)
}
