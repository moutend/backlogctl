package markdown

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/moutend/backlogctl/internal/cache"
	"github.com/moutend/backlogctl/internal/models"
	"github.com/moutend/go-backlog/pkg/types"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

type issueFrontMatter struct {
	Project   string  `yaml:"project"`
	Parent    string  `yaml:"parent"`
	IssueID   int     `yaml:"id"`
	IssueType string  `yaml:"type"`
	IssueKey  string  `yaml:"issue"`
	Priority  string  `yaml:"priority"`
	Status    string  `yaml:"status"`
	Assignee  string  `yaml:"assignee"`
	Start     string  `yaml:"start"`
	Due       string  `yaml:"due"`
	Estimated float64 `yaml:"estimated"`
	Actual    float64 `yaml:"actual"`
	Summary   string  `yaml:"summary"`
}

func ReadIssue(ctx context.Context, data []byte) (*types.Issue, error) {
	var fo issueFrontMatter

	n := bytes.Index(data[4:], []byte("---"))

	if n < 0 {
		return nil, fmt.Errorf("invalid markdown")
	}
	if err := yaml.Unmarshal(data[4:4+n], &fo); err != nil {
		return nil, fmt.Errorf("failed to parse markdown: %w", err)
	}

	var projectID int64
	issue := &types.Issue{}

	if fo.IssueKey != "" && !strings.HasSuffix(fo.IssueKey, "-0") {
		if err := cache.FetchIssue(ctx, fo.IssueKey); err != nil {
			return nil, fmt.Errorf("failed to fetch issue %q: %w", fo.IssueKey, err)
		}

		var err error

		issue, err = cache.LoadIssue(ctx, fo.IssueKey)

		if err != nil {
			return nil, fmt.Errorf("issue %q not found: %w", fo.IssueKey, err)
		}
	}

	issue.Summary = fo.Summary
	issue.Description = strings.TrimSuffix(string(data[8+n:]), "\n")

	if fo.Project != "" {
		if err := cache.FetchProjects(ctx); err != nil {
			return nil, fmt.Errorf("failed to fetch projects: %w", err)
		}

		p, err := models.Projects(
			models.ProjectWhere.ProjectKey.EQ(fo.Project),
		).One(ctx, boil.GetContextDB())

		if err != nil {
			return nil, fmt.Errorf("invalid project %q: %w", fo.Project, err)
		}

		projectID = p.ID

		issue.ProjectId = &projectID
	}
	if fo.Parent != "" {
		if err := cache.FetchIssue(ctx, fo.Parent); err != nil {
			return nil, fmt.Errorf("failed to fetch issue %q: %w", fo.Parent, err)
		}

		i, err := models.Issues(
			models.IssueWhere.IssueKey.EQ(fo.Parent),
		).One(ctx, boil.GetContextDB())

		if err != nil {
			return nil, fmt.Errorf("invalid issue %q: %w", fo.Parent, err)
		}

		parentIssueId := i.ID

		issue.ParentIssueId = &parentIssueId
	}
	if fo.IssueType != "" {
		if err := cache.FetchIssueTypes(ctx, fmt.Sprint(projectID)); err != nil {
			return nil, fmt.Errorf("failed to fetch issue type")
		}

		it, err := models.IssueTypes(
			models.IssueTypeWhere.Name.EQ(fo.IssueType),
			models.IssueTypeWhere.ProjectID.EQ(projectID),
		).One(ctx, boil.GetContextDB())

		if err != nil {
			return nil, fmt.Errorf("invalid issue type: %q", fo.IssueType)
		}

		issue.IssueType = cache.TypesIssueType(it)
	}
	if fo.Status != "" {
		if err := cache.FetchProjectStatuses(ctx, fmt.Sprint(projectID)); err != nil {
			return nil, fmt.Errorf("failed to fetch project statuses")
		}

		ps, err := models.ProjectStatuses(
			models.ProjectStatusWhere.Name.EQ(fo.Status),
			models.ProjectStatusWhere.ProjectID.EQ(projectID),
		).One(ctx, boil.GetContextDB())

		if err != nil {
			return nil, fmt.Errorf("invalid status: %q", fo.Status)
		}

		issue.Status = cache.TypesProjectStatus(ps)
	}
	if fo.Priority != "" {
		if err := cache.FetchPriorities(ctx); err != nil {
			return nil, fmt.Errorf("failed to fetch priorities")
		}

		p, err := models.Priorities(
			models.PriorityWhere.Name.EQ(fo.Priority),
		).One(ctx, boil.GetContextDB())

		if err != nil {
			return nil, fmt.Errorf("invalid priority: %q", fo.Priority)
		}

		issue.Priority = cache.TypesPriority(p)
	}
	if fo.Start != "" {
		t, err := time.Parse("2006-01-02", fo.Start)

		if err != nil {
			return nil, fmt.Errorf("invalid start date: %q", fo.Start)
		}

		issue.StartDate = types.NewDate(t.Format(time.RFC3339))
	}
	if fo.Due != "" {
		t, err := time.Parse("2006-01-02", fo.Due)

		if err != nil {
			return nil, fmt.Errorf("invalid due date: %q", fo.Due)
		}

		issue.DueDate = types.NewDate(t.Format(time.RFC3339))
	}
	if fo.Estimated > 0 {
		issue.EstimatedHours = types.NewHours(fo.Estimated)
	}
	if fo.Actual > 0 {
		issue.ActualHours = types.NewHours(fo.Actual)
	}

	return issue, nil
}

func WriteIssue(ctx context.Context, output io.Writer, issueIdOrKey string) error {
	if err := cache.FetchProjects(ctx); err != nil {
		return fmt.Errorf("failed to fetch projects: %w", err)
	}
	if err := cache.FetchIssue(ctx, issueIdOrKey); err != nil {
		return fmt.Errorf("failed to fetch issue: %w", err)
	}

	var result struct {
		ProjectName          string       `boil:"project_name"`
		ProjectKey           string       `boil:"project_key"`
		ParentIssueID        null.Int64   `boil:"parent_issue_id"`
		IssueID              int64        `boil:"issue_id"`
		IssueSummary         string       `boil:"issue_summary"`
		IssueDescription     string       `boil:"issue_description"`
		IssueKey             string       `boil:"issue_key"`
		IssueType            string       `boil:"issue_type"`
		IssueStatus          string       `boil:"issue_status"`
		IssuePriority        string       `json:"issue_priority"`
		IssueEstimatedHours  null.Float64 `boil:"issue_estimated_hours"`
		IssueActualHours     null.Float64 `boil:"issue_actual_hours"`
		IssueStartDate       null.Time    `boil:"issue_start_date"`
		IssueDueDate         null.Time    `boil:"issue_due_date"`
		IssueCreatedAt       null.Time    `boil:"issue_created_at"`
		IssueUpdatedAt       null.Time    `boil:"issue_updated_at"`
		IssueCreatedUserName null.String  `boil:"issue_created_user_name"`
		IssueUpdatedUserName null.String  `boil:"issue_updated_user_name"`
		IssueAssigneeName    null.String  `boil:"issue_assignee_name"`
	}

	query := `
SELECT
  p.name AS project_name
, p.project_key AS project_key
, i.parent_issue_id AS parent_issue_id
, i.id AS issue_id
, i.summary AS issue_summary
, i.description AS issue_description
, i.issue_key AS issue_key
, it.name AS issue_type
, ps.name AS issue_status
, pr.name AS issue_priority
, i.estimated_hours AS issue_estimated_hours
, i.actual_hours AS issue_actual_hours
, i.start_date AS issue_start_date
, i.due_date AS issue_due_date
, i.created_at AS issue_created_at
, i.updated_at AS issue_updated_at
, c.name AS issue_created_user_name
, u.name AS issue_updated_user_name
, a.name AS issue_assignee_name
FROM issue i
INNER JOIN project p ON p.id = i.project_id
INNER JOIN project_status ps ON ps.id = i.status_id
INNER JOIN priority pr ON pr.id = i.priority_id
INNER JOIN issue_type it ON it.id = i.issue_type_id
LEFT JOIN user a ON a.id = i.assignee_id
LEFT JOIN user c ON c.id = i.created_user_id
LEFT JOIN user u ON u.id = i.updated_user_id
WHERE i.id = ?
OR i.issue_key = ?
`

	if err := queries.Raw(query, issueIdOrKey, issueIdOrKey).Bind(ctx, boil.GetContextDB(), &result); err != nil {
		return fmt.Errorf("failed to read cache: %w", err)
	}

	var parent *models.Issue

	if result.ParentIssueID.Valid {
		if err := cache.FetchIssue(ctx, fmt.Sprint(result.ParentIssueID.Int64)); err != nil {
			return fmt.Errorf("failed to fetch parent issue %d: %w", result.ParentIssueID, err)
		}

		var err error

		if parent, err = models.FindIssue(ctx, boil.GetContextDB(), result.ParentIssueID.Int64); err != nil {
			return fmt.Errorf("parent issue %d not found: %w", result.ParentIssueID, err)
		}
	}

	fmt.Fprintln(output, "---")
	fmt.Fprintf(output, "summary: %s\n", result.IssueSummary)
	fmt.Fprintf(output, "issue: %s\n", result.IssueKey)

	if parent != nil {
		fmt.Fprintf(output, "parent: %s\n", parent.IssueKey)
	}

	fmt.Fprintf(output, "project: %s\n", result.ProjectKey)
	fmt.Fprintf(output, "type: %s\n", result.IssueType)
	fmt.Fprintf(output, "status: %s\n", result.IssueStatus)
	fmt.Fprintf(output, "priority: %s\n", result.IssuePriority)

	if result.IssueStartDate.Valid {
		fmt.Fprintf(output, "start: %s\n", result.IssueStartDate.Time.UTC().Format("2006-01-02"))
	}
	if result.IssueDueDate.Valid {
		fmt.Fprintf(output, "due: %s\n", result.IssueDueDate.Time.UTC().Format("2006-01-02"))
	}
	if result.IssueEstimatedHours.Valid {
		fmt.Fprintf(output, "estimated: %.1f\n", result.IssueEstimatedHours.Float64)
	}
	if result.IssueActualHours.Valid {
		fmt.Fprintf(output, "actual: %.1f\n", result.IssueActualHours.Float64)
	}
	if result.IssueAssigneeName.Valid {
		fmt.Fprintf(output, "assignee: %s\n", result.IssueAssigneeName.String)
	}
	if result.IssueCreatedAt.Valid {
		fmt.Fprintf(output, "created: %s by %s\n", result.IssueCreatedAt.Time.Format("2006-01-02"), result.IssueCreatedUserName.String)
	}
	if result.IssueUpdatedAt.Valid {
		fmt.Fprintf(output, "updated: %s by %s\n", result.IssueUpdatedAt.Time.Format("2006-01-02"), result.IssueUpdatedUserName.String)
	}

	fmt.Fprintf(output, "url: https://%s/view/%s\n", cache.BacklogSpace, result.IssueKey)
	fmt.Fprintln(output, "---")
	fmt.Fprint(output, result.IssueDescription)

	return nil
}
