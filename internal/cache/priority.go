package cache

import (
	"context"
	"database/sql"
	"time"

	"github.com/moutend/backlogctl/internal/models"
	"github.com/moutend/go-backlog/pkg/types"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TypesPriority(a *models.Priority) *types.Priority {
	if a == nil {
		return nil
	}

	return &types.Priority{
		Id:   a.ID,
		Name: a.Name,
	}
}

func ModelsPriority(a *models.Priority, b *types.Priority) *models.Priority {
	now := time.Now().UTC()

	if a == nil {
		a = &models.Priority{
			CreatedAt: now,
			UpdatedAt: now,
		}
	} else {
		a.UpdatedAt = now
	}

	a.ID = b.Id
	a.Name = b.Name

	return a
}

func SavePriority(ctx context.Context, priority *types.Priority) error {
	if priority == nil {
		return nil
	}

	p, err := models.FindPriority(ctx, boil.GetContextDB(), priority.Id)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		err = ModelsPriority(p, priority).Insert(ctx, boil.GetContextDB(), boil.Infer())
	} else {
		_, err = ModelsPriority(p, priority).Update(ctx, boil.GetContextDB(), boil.Infer())
	}

	return err
}

func FetchPriorities(ctx context.Context) error {
	ps, err := backlog.GetPrioritiesContext(ctx)

	if err != nil {
		return err
	}
	for _, p := range ps {
		if err := SavePriority(ctx, p); err != nil {
			return err
		}
	}

	return nil
}
