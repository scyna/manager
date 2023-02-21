package service

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/domain"
	"github.com/scyna/manager/model"
	proto "github.com/scyna/manager/proto/generated"
)

func CreateModuleHandler(ctx *scyna.Endpoint, request *proto.CreateModuleRequest) scyna.Error {
	service := domain.NewModuleService(&ctx.Context)

	if err := service.AssureModuleNotExist(request.Code); err == nil {
		return model.MODULE_EXISTED
	}

	module, err := model.NewModule(request.Code, request.Secret)
	if err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	// Because module has no entity_id => it creates directly module in repostory
	if err := service.Repository.CreateModule(&module); err != nil {
		return err
	}

	return scyna.OK
}
