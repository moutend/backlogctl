package cache

import (
	"context"
	"database/sql"
	"time"

	"github.com/moutend/backlogctl/internal/models"
	"github.com/moutend/go-backlog/pkg/types"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func ModelsProject(a *models.Project, b *types.Project) *models.Project {
	now := time.Now().UTC()

	if a == nil {
		a = &models.Project{
			CreatedAt: now,
			UpdatedAt: now,
		}
	} else {
		a.UpdatedAt = now
	}

	a.ID = b.Id
	a.ProjectKey = b.ProjectKey
	a.Name = b.Name
	a.TextFormattingRule = b.TextFormattingRule
	a.Archived = b.Archived

	return a
}

func FetchProjects(ctx context.Context) error {
	projects, err := backlog.GetProjectsContext(ctx, nil)

	if err != nil {
		return err
	}
	for _, project := range projects {
		p, err := models.FindProject(ctx, boil.GetContextDB(), int64(project.Id))

		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if err == sql.ErrNoRows {
			err = ModelsProject(p, project).Insert(ctx, boil.GetContextDB(), boil.Infer())
		} else {
			_, err = ModelsProject(p, project).Update(ctx, boil.GetContextDB(), boil.Infer())
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func TypesProjectStatus(a *models.ProjectStatus) *types.ProjectStatus {
	if a == nil {
		return nil
	}

	return &types.ProjectStatus{
		Id:           a.ID,
		ProjectId:    a.ProjectID,
		Name:         a.Name,
		Color:        a.Color,
		DisplayOrder: a.DisplayOrder,
	}
}

func ModelsProjectStatus(a *models.ProjectStatus, b *types.ProjectStatus) *models.ProjectStatus {
	now := time.Now().UTC()

	if a == nil {
		a = &models.ProjectStatus{
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
	a.DisplayOrder = b.DisplayOrder

	return a
}

func SaveProjectStatus(ctx context.Context, projectStatus *types.ProjectStatus) error {
	if projectStatus == nil {
		return nil
	}

	ps, err := models.FindProjectStatus(ctx, boil.GetContextDB(), projectStatus.Id)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		err = ModelsProjectStatus(ps, projectStatus).Insert(ctx, boil.GetContextDB(), boil.Infer())
	} else {
		_, err = ModelsProjectStatus(ps, projectStatus).Update(ctx, boil.GetContextDB(), boil.Infer())
	}

	return err
}

func FetchProjectStatuses(ctx context.Context, projectIdOrKey string) error {
	projectStatuses, err := backlog.GetProjectStatusesContext(ctx, projectIdOrKey)

	if err != nil {
		return err
	}
	for _, projectStatus := range projectStatuses {
		if err := SaveProjectStatus(ctx, projectStatus); err != nil {
			return err
		}
	}

	return nil
}
