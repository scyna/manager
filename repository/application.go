package repository

import (
	"log"

	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/model"
)

type Application struct {
	Code     string `db:"code"`
	Name     string `db:"name"`
	AuthPath string `db:"auth_url"`
}

func GetApplication(code string) (*Application, scyna.Error) {
	var ret Application
	if err := qb.Select(APPLICATION_TABLE).
		Columns("code", "name", "auth_url").
		Where(qb.Eq("code")).
		Limit(1).
		Query(scyna.DB).
		Bind(code).
		GetRelease(&ret); err != nil {
		return nil, model.APPLICATION_NOT_EXIST
	}

	return &ret, nil
}

func CreateApplication(application *Application) scyna.Error {
	if err := qb.Insert(APPLICATION_TABLE).
		Columns("code", "name", "auth_url").
		Query(scyna.DB).
		BindStruct(&application).
		ExecRelease(); err != nil {
		log.Println(err)
		return scyna.SERVER_ERROR
	}

	return nil
}

func DeleteApplication(code string) scyna.Error {
	if err := qb.Delete(APPLICATION_TABLE).
		Where(qb.Eq("code")).
		Query(scyna.DB).
		Bind(code).
		ExecRelease(); err != nil {
		return model.APPLICATION_NOT_EXIST
	}

	return nil
}

func UpdateApplication(application *Application) scyna.Error {
	if err := qb.Update(APPLICATION_TABLE).
		Set("name", "auth_url").
		Where(qb.Eq("code")).
		Query(scyna.DB).
		BindStruct(&application).
		ExecRelease(); err != nil {
		return model.APPLICATION_NOT_EXIST
	}

	return nil
}

func ListApplication() ([]*Application, scyna.Error) {
	var ret []*Application
	if err := qb.Select(APPLICATION_TABLE).
		Columns("code", "name", "auth_url").
		Query(scyna.DB).
		SelectRelease(&ret); err != nil {
		return nil, scyna.SERVER_ERROR
	}

	return ret, nil
}
