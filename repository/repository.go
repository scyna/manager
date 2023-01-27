package repository

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/domain"
)

const MODULE_TABLE = "scyna.module"

type managerRepository struct {
	LOG scyna.Logger
}

func NewRepository(LOG scyna.Logger) domain.IRepository {
	return &managerRepository{LOG: LOG}
}
