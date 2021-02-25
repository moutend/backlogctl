package app

import (
	"strconv"

	"github.com/moutend/backlogctl/internal/cache"
	"github.com/moutend/backlogctl/internal/markdown"
	"github.com/moutend/backlogctl/internal/models"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var wikiCommand = &cobra.Command{
	Use:     "wiki",
	Aliases: []string{"w"},
	Short:   "provides CRUD operations",
	RunE:    wikiCommandRunE,
}

func wikiCommandRunE(cmd *cobra.Command, args []string) error {
	return nil
}

var wikiListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "list wikis",
	RunE:    wikiListCommandRunE,
}

func wikiListCommandRunE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	db := boil.GetContextDB()

	if err := cache.FetchProjects(cmd.Context()); err != nil {
		return err
	}

	projects, err := models.Projects().All(ctx, db)

	if err != nil {
		return err
	}
	for _, project := range projects {
		if err := cache.FetchWikis(ctx, project.ProjectKey); err != nil {
			return err
		}
	}

	wikis, err := models.Wikis().All(ctx, db)

	if err != nil {
		return err
	}

	for _, wiki := range wikis {
		cmd.Printf("- %s (%d)\n", wiki.Name, wiki.ID)
	}

	return nil
}

var wikiReadCommand = &cobra.Command{
	Use:     "read",
	Aliases: []string{"r"},
	Short:   "read wiki",
	RunE:    wikiReadCommandRunE,
}

func wikiReadCommandRunE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if len(args) < 1 {
		return nil
	}

	wikiID, err := strconv.ParseInt(args[0], 10, 64)

	if err != nil {
		return err
	}
	if err := markdown.WriteWiki(ctx, cmd.OutOrStdout(), int64(wikiID)); err != nil {
		return err
	}

	return nil
}

func init() {
	RootCommand.AddCommand(wikiCommand)

	wikiCommand.AddCommand(wikiListCommand)
	wikiCommand.AddCommand(wikiReadCommand)
}
