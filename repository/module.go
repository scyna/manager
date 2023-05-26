package repository

import (
	"encoding/json"
	"log"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
	scyna_const "github.com/scyna/core/const"
	"github.com/scyna/manager/model"
)

type Module struct {
	Code   string `db:"code"`
	Name   string `db:"name"`
	Secret string `db:"secret"`
}

type Config struct {
	NatsURL      string `json:"NatsUrl"`
	NatsUsername string `json:"NatsUsername"`
	NatsPassword string `json:"NatsPassword"`
	DBHost       string `json:"DBHost"`
	DBUsername   string `json:"DBUsername"`
	DBPassword   string `json:"DBPassword"`
	DBLocation   string `json:"DBLocation"`
}

func GetModule(code string) (*Module, scyna.Error) {
	var ret Module

	if err := qb.Select(MODULE_TABLE).
		Columns("code", "name").
		Where(qb.Eq("code")).
		Limit(1).
		Query(scyna.DB).
		Bind(code).
		GetRelease(&ret); err != nil {
		return nil, model.MODULE_NOT_EXIST
	}

	return &ret, nil
}

func CreateModule(module *Module, config string) scyna.Error {

	session := scyna.DB.Session
	batch := session.NewBatch(gocql.LoggedBatch)
	batch.Query("INSERT INTO "+MODULE_TABLE+" (code, secret, name) VALUES (?,?,?);",
		module.Code, module.Secret, module.Name)
	batch.Query("INSERT INTO "+SETTING_TABLE+" (module, key, value) VALUES (?,?,?);",
		module.Code, scyna_const.SETTING_KEY, config)
	if err := session.ExecuteBatch(batch); err != nil {
		log.Print("Error:", err)
		return scyna.SERVER_ERROR
	}

	return nil
}

func ChangePasswordModule(code string, secret string) scyna.Error {
	if err := qb.Update(MODULE_TABLE).
		Set("secret").
		Where(qb.Eq("code")).
		Query(scyna.DB).
		Bind(secret, code).
		ExecRelease(); err != nil {
		return model.MODULE_NOT_EXIST
	}

	return nil
}

func DeleteModule(code string) scyna.Error {
	if err := qb.Delete(MODULE_TABLE).
		Where(qb.Eq("code")).
		Query(scyna.DB).
		Bind(code).
		ExecRelease(); err != nil {
		return model.MODULE_NOT_EXIST
	}

	return nil
}

func ListModule() ([]*Module, scyna.Error) {
	var modules []*Module
	if err := qb.Select(MODULE_TABLE).
		Columns("code", "name").
		Query(scyna.DB).
		SelectRelease(&modules); err != nil {
		log.Print(err)
		return nil, scyna.SERVER_ERROR
	}
	return modules, nil
}

func UpdateModule(module *Module) scyna.Error {
	if err := qb.Update(MODULE_TABLE).
		Set("name").
		Where(qb.Eq("code")).
		Query(scyna.DB).
		Bind(module.Name, module.Code).
		ExecRelease(); err != nil {
		return model.MODULE_NOT_EXIST
	}

	return nil
}

func UpdateConfigModule(code string, config string) scyna.Error {
	if err := qb.Update(SETTING_TABLE).
		Set("value").
		Where(qb.Eq("module"), qb.Eq("key")).
		Query(scyna.DB).
		Bind(config, code, scyna_const.SETTING_KEY).
		ExecRelease(); err != nil {
		return model.MODULE_NOT_EXIST
	}

	return nil
}

func GetConfigModule(code string) (*Config, scyna.Error) {
	var ret string

	if err := qb.Select(SETTING_TABLE).
		Columns("value").
		Where(qb.Eq("module"), qb.Eq("key")).
		Limit(1).
		Query(scyna.DB).
		Bind(code, scyna_const.SETTING_KEY).
		GetRelease(&ret); err != nil {
		return nil, model.MODULE_NOT_EXIST
	}

	var config Config

	if err := json.Unmarshal([]byte(ret), &config); err != nil {
		log.Print(err)
		return nil, scyna.SERVER_ERROR
	}

	return &config, nil
}
