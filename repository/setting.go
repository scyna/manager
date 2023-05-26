package repository

import (
	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/model"
)

type Setting struct {
	Module string `db:"module"`
	Key    string `db:"key"`
	Value  string `db:"value"`
}

func GetSetting(module string, key string) (*Setting, scyna.Error) {
	var ret Setting
	if err := qb.Select(SETTING_TABLE).
		Columns("value", "module", "key").
		Where(qb.Eq("module"), qb.Eq("key")).
		Limit(1).
		Query(scyna.DB).
		Bind(module, key).
		GetRelease(&ret); err != nil {
		return nil, model.SETTING_NOT_EXIST
	}
	return &ret, nil
}

func ListSetting(module string) ([]Setting, scyna.Error) {
	var ret []Setting
	if err := qb.Select(SETTING_TABLE).
		Columns("key", "module", "value").
		Where(qb.Eq("module")).
		Query(scyna.DB).
		Bind(module).
		SelectRelease(&ret); err != nil {
		return nil, scyna.SERVER_ERROR
	}
	return ret, nil
}

func UpdateSetting(setting *Setting) scyna.Error {
	if err := qb.Update(SETTING_TABLE).
		Set("value").
		Where(qb.Eq("module"), qb.Eq("key")).
		Query(scyna.DB).
		BindStruct(&setting).
		ExecRelease(); err != nil {
		return scyna.SERVER_ERROR
	}
	return nil
}

func CreateSetting(setting *Setting) scyna.Error {
	if err := qb.Insert(SETTING_TABLE).
		Columns("module", "key", "value").
		Query(scyna.DB).
		BindStruct(&setting).
		ExecRelease(); err != nil {
		return scyna.SERVER_ERROR
	}
	return nil
}

func DeleteSetting(module string, key string) scyna.Error {
	if err := qb.Delete(SETTING_TABLE).
		Where(qb.Eq("module"), qb.Eq("key")).
		Query(scyna.DB).
		Bind(module, key).
		ExecRelease(); err != nil {
		return scyna.SERVER_ERROR
	}
	return nil
}
