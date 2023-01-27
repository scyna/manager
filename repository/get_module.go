package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/model"
)

func (r *managerRepository) GetModule(code string) (*model.Module, scyna.Error) {
	var module model.Module

	if err := qb.Select(MODULE_TABLE).
		Columns("code").
		Where(qb.Eq("code")).
		Limit(1).
		Query(scyna.DB).Bind(code).GetRelease(&module); err != nil {
		r.LOG.Error(err.Error())
		return nil, model.MODULE_NOT_EXIST
	}

	ret := &model.Module{
		Code: module.Code,
	}

	return ret, nil
}
