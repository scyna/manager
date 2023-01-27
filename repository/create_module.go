package repository

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/model"
)

func (r *Repository) CreateModule(module *model.Module) scyna.Error {

	// if err := qb.Insert("scyna.module").
	// 	Columns("code", "secret").
	// 	Query(scyna.DB).
	// 	BindStruct(context).
	// 	ExecRelease(); err != nil {
	// 	return scyna.SERVER_ERROR
	// }

	return nil
}
