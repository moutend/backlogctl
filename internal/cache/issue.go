package cache

import (
	"context"
	"database/sql"
	"net/url"
	"time"

	"github.com/moutend/go-backlog/pkg/types"

	"github.com/moutend/backlogctl/internal/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TypesIssueType(a *models.IssueType) *types.IssueType {
	if a == nil {
		return nil
	}

	return &types.IssueType{
		Id:           a.ID,
		ProjectId:    a.ProjectID,
		Name:         a.Name,
		Color:        a.Color,
		DisplayOrder: int(a.DisplayOrder),
	}
}

func ModelsIssueType(a *models.IssueType, b *types.IssueType) *models.IssueType {
	now := time.Now().UTC()

	if a == nil {
		a = &models.IssueType{
			CreatedAt: now,
			UpdatedAt: now,
		}
	} else {
		a.UpdatedAt = now
	}

	a.ID = b.Id
	a.ProjectID = b.ProjectId
	a.Name = b.Name
	a.Color = b.Color
	a.DisplayOrder = int64(b.DisplayOrder)

	return a
}

func SaveIssueType(ctx context.Context, issueType *types.IssueType) error {
	it, err := models.FindIssueType(ctx, boil.GetContextDB(), issueType.Id)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		err = ModelsIssueType(it, issueType).Insert(ctx, boil.GetContextDB(), boil.Infer())
	} else {
		_, err = ModelsIssueType(it, issueType).Update(ctx, boil.GetContextDB(), boil.Infer())
	}

	return err
}

func FetchIssueTypes(ctx context.Context, projectIdOrKey string) error {
	issueTypes, err := backlog.GetIssueTypesContext(ctx, projectIdOrKey)

	if err != nil {
		return err
	}
	for _, issueType := range issueTypes {
		if err := SaveIssueType(ctx, issueType); err != nil {
			return err
		}
	}

	return nil
}

func TypesIssue(a *models.Issue) *types.Issue {
	if a == nil {
		return nil
	}

	var (
		startDate      *types.Date
		dueDate        *types.Date
		estimatedHours *types.Hours
		actualHours    *types.Hours
	)

	if a.StartDate.Valid {
		startDate = types.NewDate(a.StartDate.Time.Format(time.RFC3339))
	}
	if a.DueDate.Valid {
		dueDate = types.NewDate(a.DueDate.Time.Format(time.RFC3339))
	}
	if a.EstimatedHours.Valid {
		estimatedHours = types.NewHours(a.EstimatedHours.Float64)
	}
	if a.ActualHours.Valid {
		actualHours = types.NewHours(a.ActualHours.Float64)
	}

	return &types.Issue{
		Id:             a.ID,
		ProjectId:      a.ProjectID.Ptr(),
		IssueKey:       a.IssueKey,
		KeyId:          a.KeyID,
		IssueType:      nil,
		Summary:        a.Summary,
		Description:    a.Description,
		Resolution:     nil,
		Priority:       nil,
		Status:         nil,
		Assignee:       nil,
		Category:       nil,
		Versions:       nil,
		Milestone:      nil,
		StartDate:      startDate,
		DueDate:        dueDate,
		EstimatedHours: estimatedHours,
		ActualHours:    actualHours,
		ParentIssueId:  a.ParentIssueID.Ptr(),
	}
}

func ModelsIssue(a *models.Issue, b *types.Issue) *models.Issue {
	now := time.Now()

	if a == nil {
		a = &models.Issue{
			CreatedAt: now,
			UpdatedAt: now,
		}
	} else {
		a.UpdatedAt = now
	}

	a.ID = int64(b.Id)

	if b.ProjectId != nil {
		a.ProjectID = null.Int64From(int64(*b.ProjectId))
	}

	a.IssueKey = b.IssueKey
	a.KeyID = b.KeyId

	if b.IssueType != nil {
		a.IssueTypeID = null.Int64From(b.IssueType.Id)
	}

	a.Summary = b.Summary
	a.Description = b.Description
	// a.ResolutionID   = null.Int64From(0)

	if b.Priority != nil {
		a.PriorityID = null.Int64From(b.Priority.Id)
	}
	if b.Status != nil {
		a.StatusID = null.Int64From(b.Status.Id)
	}
	if b.Assignee != nil {
		a.AssigneeID = null.Int64From(int64(b.Assignee.Id))
	}
	if b.StartDate != nil {
		a.StartDate = null.TimeFrom(b.StartDate.Time().UTC())
	}
	if b.DueDate != nil {
		a.DueDate = null.TimeFrom(b.DueDate.Time().UTC())
	}
	if b.EstimatedHours != nil {
		a.EstimatedHours = null.Float64From(float64(*b.EstimatedHours))
	}
	if b.ActualHours != nil {
		a.ActualHours = null.Float64From(float64(*b.ActualHours))
	}

	a.ParentIssueID = null.Int64FromPtr(b.ParentIssueId)

	if b.Created != nil {
		a.CreatedAt = b.Created.Time().UTC()
	}
	if b.Updated != nil {
		a.UpdatedAt = b.Updated.Time().UTC()
	}
	if b.CreatedUser != nil {
		a.CreatedUserID = b.CreatedUser.Id
	}
	if b.UpdatedUser != nil {
		a.UpdatedUserID = b.UpdatedUser.Id
	}

	return a
}

func LoadIssue(ctx context.Context, issueKeyOrId string) (*types.Issue, error) {
	var (
		p  *models.Priority
		ps *models.ProjectStatus
		a  *models.User
	)

	i, err := models.Issues(
		models.IssueWhere.IssueKey.EQ(issueKeyOrId),
	).One(ctx, boil.GetContextDB())

	if err != nil {
		return nil, err
	}

	if i.PriorityID.Valid {
		p, err = models.FindPriority(ctx, boil.GetContextDB(), i.PriorityID.Int64)

		if err != nil {
			return nil, err
		}
	}
	if i.StatusID.Valid {
		ps, err = models.FindProjectStatus(ctx, boil.GetContextDB(), i.StatusID.Int64)

		if err != nil {
			return nil, err
		}
	}
	if i.AssigneeID.Valid {
		a, err = models.FindUser(ctx, boil.GetContextDB(), i.AssigneeID.Int64)

		if err != nil {
			return nil, err
		}
	}

	issue := TypesIssue(i)

	issue.Priority = TypesPriority(p)
	issue.Status = TypesProjectStatus(ps)
	issue.Assignee = TypesUser(a)

	return issue, nil
}

func SaveIssue(ctx context.Context, issue *types.Issue) error {
	if err := SaveProjectStatus(ctx, issue.Status); err != nil {
		return err
	}
	if err := SavePriority(ctx, issue.Priority); err != nil {
		return err
	}
	if err := SaveIssueType(ctx, issue.IssueType); err != nil {
		return err
	}
	if err := SaveUser(ctx, issue.Assignee); err != nil {
		return err
	}
	if err := SaveUser(ctx, issue.CreatedUser); err != nil {
		return err
	}
	if err := SaveUser(ctx, issue.UpdatedUser); err != nil {
		return err
	}

	i, err := models.FindIssue(ctx, boil.GetContextDB(), int64(issue.Id))

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		err = ModelsIssue(i, issue).Insert(ctx, boil.GetContextDB(), boil.Infer())
	} else {
		_, err = ModelsIssue(i, issue).Update(ctx, boil.GetContextDB(), boil.Infer())
	}

	return err
}

func CreateIssue(ctx context.Context, issue *types.Issue, notifiedUsers []*types.User) (*types.Issue, error) {
	issue, err := backlog.AddIssueContext(ctx, issue, notifiedUsers)

	if err != nil {
		return nil, err
	}

	return issue, nil
}

func UpdateIssue(ctx context.Context, issue *types.Issue, notifiedUsers []*types.User, comment string) (*types.Issue, error) {
	issue, err := backlog.UpdateIssueContext(ctx, issue, notifiedUsers, comment)

	if err != nil {
		return nil, err
	}

	return issue, nil
}

func FetchIssue(ctx context.Context, issueKeyOrId string) error {
	issue, err := backlog.GetIssueContext(ctx, issueKeyOrId)

	if err != nil {
		return err
	}
	if err := SaveIssue(ctx, issue); err != nil {
		return err
	}

	return nil
}

func FetchIssues(ctx context.Context, maxIssues int, values url.Values) error {
	issues, err1 := backlog.GetAllIssuesContext(ctx, maxIssues, values)

	for _, issue := range issues {
		if err := SaveIssue(ctx, issue); err != nil {
			return err
		}
	}

	return err1
}
