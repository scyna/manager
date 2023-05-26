package repository

import (
	"log"

	"github.com/scylladb/gocqlx/v2/qb"
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/model"
)

type Client struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

func GetClient(id string) (*Client, scyna.Error) {
	var ret Client

	if err := qb.Select(CLIENT_TABLE).
		Columns("id", "name").
		Where(qb.Eq("id")).
		Limit(1).
		Query(scyna.DB).
		Bind(id).
		GetRelease(&ret); err != nil {
		return nil, model.CLIENT_NOT_EXIST
	}

	return &ret, nil
}

func ListClient() ([]Client, scyna.Error) {
	var ret []Client

	if err := qb.Select(CLIENT_TABLE).
		Columns("id", "name").
		Query(scyna.DB).
		SelectRelease(&ret); err != nil {
		return nil, scyna.SERVER_ERROR
	}

	return ret, nil
}

func CreateClient(client *Client) scyna.Error {
	if err := qb.Insert(CLIENT_TABLE).
		Columns("id", "name", "secret").
		Query(scyna.DB).
		BindStruct(&client).
		ExecRelease(); err != nil {
		log.Println(err)
		return scyna.SERVER_ERROR
	}

	return nil
}

func DeleteClient(id string) scyna.Error {
	if err := qb.Delete(CLIENT_TABLE).
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(id).
		ExecRelease(); err != nil {
		return model.CLIENT_NOT_EXIST
	}

	return nil
}

func UpdateClient(client *Client) scyna.Error {
	if err := qb.Update(CLIENT_TABLE).
		Set("name").
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(client.Name, client.ID).
		ExecRelease(); err != nil {
		return model.CLIENT_NOT_EXIST
	}

	return nil
}

func ChangeSecretClient(client *Client) scyna.Error {
	if err := qb.Update(CLIENT_TABLE).
		Set("secret").
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(client.Secret, client.ID).
		ExecRelease(); err != nil {
		return model.CLIENT_NOT_EXIST
	}

	return nil
}
