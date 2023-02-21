package domain

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/model"
	"github.com/scyna/manager/repository"
)

type moduleService struct {
	Repository *repository.ModuleRepository
	context    *scyna.Context
}

func NewModuleService(context *scyna.Context) *moduleService {
	return &moduleService{
		context:    context,
		Repository: &repository.ModuleRepository{LOG: context},
	}
}

func (s *moduleService) AssureModuleNotExist(moduleName string) scyna.Error {
	modules, err := s.Repository.ListModule()
	if err != nil {
		return model.MODULE_EXISTED
	}

	for _, module := range modules {
		if module.Code == moduleName {
			return model.MODULE_EXISTED
		}
	}

	return nil
}

func (s *moduleService) GetModule(moduleName string) (*model.Module, scyna.Error) {
	// TODO
	return &model.Module{}, nil
}
