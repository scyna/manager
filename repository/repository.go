package repository

import (
	scyna "github.com/scyna/core"
)

const MODULE_TABLE = "scyna.module"

type Repository struct {
	LOG scyna.Logger
}

func LoadRepository(LOG scyna.Logger) *Repository {
	return &Repository{LOG: LOG}
}
