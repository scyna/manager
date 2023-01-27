package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/model"
)

func (r *managerRepository) CreateModule(LOG scyna.Logger, context *model.Module) scyna.Error {

	if err := qb.Insert("scyna.module").
		Columns("code", "secret").
		Query(scyna.DB).
		BindStruct(context).
		ExecRelease(); err != nil {
		return scyna.SERVER_ERROR
	}

	return nil
}
