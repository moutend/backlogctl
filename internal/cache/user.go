package cache

import (
	"context"
	"database/sql"
	"time"

	"github.com/moutend/backlogctl/internal/models"
	"github.com/moutend/go-backlog/pkg/types"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TypesUser(a *models.User) *types.User {
	if a == nil {
		return nil
	}

	return &types.User{
		Id:          a.ID,
		UserId:      a.UserID,
		Name:        a.Name,
		RoleType:    int(a.RoleType),
		Lang:        a.Language,
		MailAddress: a.Email,
	}
}

func ModelsUser(a *models.User, b *types.User) *models.User {
	now := time.Now().UTC()

	if a == nil {
		a = &models.User{
			CreatedAt: now,
			UpdatedAt: now,
		}
	} else {
		a.UpdatedAt = now
	}

	a.ID = b.Id
	a.UserID = b.UserId
	a.Name = b.Name
	a.RoleType = int64(b.RoleType)
	a.Language = b.Lang
	a.Email = b.MailAddress

	return a
}

func SaveUser(ctx context.Context, user *types.User) error {
	if user == nil {
		return nil
	}

	u, err := models.FindUser(ctx, boil.GetContextDB(), user.Id)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		err = ModelsUser(u, user).Insert(ctx, boil.GetContextDB(), boil.Infer())
	} else {
		_, err = ModelsUser(u, user).Update(ctx, boil.GetContextDB(), boil.Infer())
	}

	return err
}

func FetchMyself(ctx context.Context) error {
	myself, err := backlog.GetMyselfContext(ctx)

	if err != nil {
		return err
	}

	u, err := models.FindUser(ctx, boil.GetContextDB(), int64(myself.Id))

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	u = ModelsUser(u, myself)

	u.IsMyself = true

	if err == sql.ErrNoRows {
		err = u.Insert(ctx, boil.GetContextDB(), boil.Infer())
	} else {
		_, err = u.Update(ctx, boil.GetContextDB(), boil.Infer())
	}

	return err
}
