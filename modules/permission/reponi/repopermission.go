package reponi

import (
	"002/modules/permission/model"
	"database/sql"
)

type repoPermission struct {
	db *sql.DB
}

type RepoPermission interface {
	InsertPermission(req model.Permission) error
	SelectAll() ([]model.Permission, error)
	DeleteId(id int64) error
}

func NewRepoPermission(db *sql.DB) *repoPermission {
	return &repoPermission{
		db: db,
	}
}
