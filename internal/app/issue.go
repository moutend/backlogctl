package app

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/moutend/backlogctl/internal/cache"
	"github.com/moutend/backlogctl/internal/markdown"
	"github.com/moutend/backlogctl/internal/models"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

var issueCommand = &cobra.Command{
	Use:     "issue",
	Aliases: []string{"i"},
	RunE:    issueCommandRunE,
}

func issueCommandRunE(cmd *cobra.Command, args []string) error {
	return nil
}

var issueListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	RunE:    issueListCommandRunE,
}

func issueListCommandRunE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	db := boil.GetContextDB()

	var cancel context.CancelFunc

	if timeout, _ := cmd.Flags().GetDuration("timeout"); timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)

		defer cancel()
	}
	if err := cache.FetchProjects(cmd.Context()); err != nil {
		return err
	}
	if err := cache.FetchMyself(cmd.Context()); err != nil {
		return err
	}

	values := url.Values{}

	values.Add("order", "desc")

	if yes, _ := cmd.Flags().GetBool("myself"); yes {
		myself, err := models.Users(
			models.UserWhere.IsMyself.EQ(true),
		).One(cmd.Context(), boil.GetContextDB())

		if err != nil {
			return err
		}

		values.Add("assigneeId[]", fmt.Sprint(myself.ID))
	}
	if projectKey, _ := cmd.Flags().GetString("project"); projectKey != "" {
		projects, err := models.Projects().All(ctx, db)

		if err != nil {
			return err
		}
		for _, project := range projects {
			if projectKey == project.ProjectKey {
				values.Add("projectId", fmt.Sprint(project.ID))
			}
		}
	}

	values.Add("sort", "created")

	maxIssues, _ := cmd.Flags().GetInt("max")

	if yes, _ := cmd.Flags().GetBool("all"); yes {
		maxIssues = -1
	}
	if err := cache.FetchIssues(ctx, maxIssues, values); err != nil {
		return err
	}

	var issues []*struct {
		IssueKey string `boil:"issue_key"`
		Summary  string `boil:"summary"`
		Status   string `boil:"status"`
	}

	query := `
SELECT
  i.id AS id
, i.issue_key AS issue_key
, i.summary AS summary
, COALESCE(ps.name, '') AS status
FROM issue i
LEFT JOIN project_status ps ON ps.id = i.status_id
ORDER BY i.created_at DESC
`

	if err := queries.Raw(query).Bind(ctx, boil.GetContextDB(), &issues); err != nil {
		return err
	}
	for _, issue := range issues {
		cmd.Printf("- %s (%s) %s\n", issue.IssueKey, issue.Status, issue.Summary)
	}

	return nil
}

var issueCreateCommand = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	RunE:    issueCreateCommandRunE,
}

func issueCreateCommandRunE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if len(args) < 1 {
		return nil
	}

	data, err := ioutil.ReadFile(args[0])

	if err != nil {
		return err
	}

	issue, err := markdown.ReadIssue(ctx, data)

	if err != nil {
		return err
	}

	issue, err = cache.CreateIssue(ctx, issue, nil)

	if err != nil {
		return err
	}

	cmd.Println("created", issue.IssueKey)

	return nil
}

var issueReadCommand = &cobra.Command{
	Use:     "read",
	Aliases: []string{"r"},
	RunE:    issueReadCommandRunE,
}

func issueReadCommandRunE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if len(args) < 1 {
		return nil
	}
	if err := markdown.WriteIssue(ctx, cmd.OutOrStdout(), args[0]); err != nil {
		return err
	}

	return nil
}

var issueUpdateCommand = &cobra.Command{
	Use:     "update",
	Aliases: []string{"u"},
	RunE:    issueUpdateCommandRunE,
}

func issueUpdateCommandRunE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if len(args) < 1 {
		return nil
	}

	data, err := ioutil.ReadFile(args[0])

	if err != nil {
		return err
	}

	issue, err := markdown.ReadIssue(ctx, data)

	if err != nil {
		return err
	}

	comment, _ := cmd.Flags().GetString("comment")

	if _, err := cache.UpdateIssue(ctx, issue, nil, comment); err != nil {
		return err
	}

	cmd.Println("updated", issue.IssueKey)

	return nil
}

func init() {
	RootCommand.AddCommand(issueCommand)

	issueListCommand.Flags().BoolP("all", "a", false, "Fetch all issues")
	issueListCommand.Flags().IntP("max", "", 20, "maximum issues to fetch")
	issueListCommand.Flags().BoolP("desc", "", true, "Print issues descending order")
	issueListCommand.Flags().BoolP("myself", "m", false, "Select issues which assigned to myself")
	issueListCommand.Flags().StringP("project", "", "", "Filtered by project key")
	issueListCommand.Flags().StringP("priority", "", "", "Specify issue priority")

	issueCommand.AddCommand(issueListCommand)
	issueCommand.AddCommand(issueCreateCommand)
	issueCommand.AddCommand(issueReadCommand)
	issueCommand.AddCommand(issueUpdateCommand)
}
