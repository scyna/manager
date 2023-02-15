package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/model"
)

func (r *ModuleRepository) GetModule(moduleName string) (*model.Module, scyna.Error) {
	var module struct {
		code   string `db:"code"`
		secret string `db:"secret"`
	}

	if err := qb.Select(MODULE_TABLE).
		Columns("code", "secret").
		Where(qb.Eq("code")).
		Limit(1).
		Query(scyna.DB).
		Bind(moduleName).
		GetRelease(&module); err != nil {
		r.LOG.Error(err.Error())
		return nil, model.MODULE_NOT_EXIST
	}

	var ret model.Module
	if moduleData, err := model.NewModule(module.code, module.secret); err != nil {
		r.LOG.Error(err.Error())
		return nil, model.MODULE_NOT_EXIST
	} else {
		ret = moduleData
	}

	return &ret, nil
}
