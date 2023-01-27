package domain

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/model"
)

type IRepository interface {
	GetModule(code string) (*model.Module, scyna.Error)
	CreateModule(LOG scyna.Logger, account *model.Module) scyna.Error
}

type RepositoryCreator func(LOG scyna.Logger) IRepository

var repositoryCreator RepositoryCreator

func LoadRepository(LOG scyna.Logger) IRepository {
	if repositoryCreator == nil {
		panic("No RepositoryCreator attached")
	}
	return repositoryCreator(LOG)
}

func AttachRepositoryCreator(rc RepositoryCreator) {
	repositoryCreator = rc
}
