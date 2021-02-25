package cache

import (
	"context"
	"database/sql"
	"time"

	"github.com/moutend/backlogctl/internal/models"
	"github.com/moutend/go-backlog/pkg/types"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func ModelsWiki(a *models.Wiki, b *types.Wiki) *models.Wiki {
	now := time.Now().UTC()

	if a == nil {
		a = &models.Wiki{
			CreatedAt: now,
			UpdatedAt: now,
		}
	} else {
		a.UpdatedAt = now
	}

	a.ID = b.Id
	a.ProjectID = b.ProjectId
	a.Name = b.Name
	a.Content = b.Content

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

func SaveWiki(ctx context.Context, wiki *types.Wiki) error {
	w, err := models.FindWiki(ctx, boil.GetContextDB(), wiki.Id)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		err = ModelsWiki(w, wiki).Insert(ctx, boil.GetContextDB(), boil.Infer())
	} else {
		_, err = ModelsWiki(w, wiki).Update(ctx, boil.GetContextDB(), boil.Infer())
	}

	return err
}

func FetchWikis(ctx context.Context, projectIDOrKey string) error {
	wikis, err := backlog.GetWikisContext(ctx, projectIDOrKey, nil)

	if err != nil {
		return err
	}
	for _, wiki := range wikis {
		if err := SaveWiki(ctx, wiki); err != nil {
			return err
		}
	}

	return nil
}

func FetchWiki(ctx context.Context, wikiID int64) error {
	wiki, err := backlog.GetWikiContext(ctx, uint64(wikiID))

	if err != nil {
		return err
	}
	if err := SaveWiki(ctx, wiki); err != nil {
		return err
	}

	return nil
}
