package markdown

import (
	"context"
	"fmt"
	"io"

	"github.com/moutend/backlogctl/internal/cache"
	"github.com/moutend/backlogctl/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func WriteWiki(ctx context.Context, output io.Writer, wikiID int64) error {
	if err := cache.FetchWiki(ctx, wikiID); err != nil {
		return err
	}

	wiki, err := models.FindWiki(ctx, boil.GetContextDB(), wikiID)

	if err != nil {
		return err
	}

	project, err := models.FindProject(ctx, boil.GetContextDB(), wiki.ProjectID)

	if err != nil {
		return err
	}

	fmt.Fprintln(output, "---")
	fmt.Fprintf(output, "name: %s\n", wiki.Name)
	fmt.Fprintf(output, "project: %s\n", project.ProjectKey)
	fmt.Fprintf(output, "wiki: %d\n", wiki.ID)
	fmt.Fprintf(output, "url: https://%s/alias/wiki/%d\n", cache.BacklogSpace, wiki.ID)
	fmt.Fprintln(output, "---")
	fmt.Fprint(output, wiki.Content)

	return nil
}
