package biz

import (
	"002/modules/permission/model"
	"002/modules/permission/reponi"
)

type bizPermission struct {
	repoPermission reponi.RepoPermission
}

type BizPermission interface {
	InsertPermission(req model.Permission) error
	SelectAll() ([]model.Permission, error)
	DeleteId(id int64) error
}

func NewBizPermission(repo reponi.RepoPermission) *bizPermission {
	return &bizPermission{
		repoPermission: repo,
	}
}

func (b *bizPermission) InsertPermission(req model.Permission) error {
	return b.repoPermission.InsertPermission(req)
}

func (b *bizPermission) SelectAll() ([]model.Permission, error) {
	return b.repoPermission.SelectAll()
}

func (b *bizPermission) DeleteId(id int64) error {
	return b.repoPermission.DeleteId(id)
}
