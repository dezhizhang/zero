package dao

import (
	"context"
	"user/model"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (d *UserDao) Save(ctx context.Context, user *model.User) error {
	return nil
}
