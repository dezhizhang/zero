package repo

import (
	"context"
	"user/model"
)

type UserRepo interface {
	Save(ctx context.Context, user *model.User) error
}
