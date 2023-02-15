package service

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/domain"
	"github.com/scyna/manager/model"
	proto "github.com/scyna/manager/proto/generated"
)

func CreateModuleHandler(ctx *scyna.Endpoint, request *proto.CreateModuleRequest) scyna.Error {
	service := domain.NewModuleService(&ctx.Context)

	if _, err := service.Repository.GetModule(request.Code); err == nil {
		return model.MODULE_EXISTED
	}

	module, err := model.NewModule(request.Code, request.Secret)
	if err != nil {
		ctx.Logger.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	if err := service.Repository.CreateModule(&module); err != nil {
		return err
	}

	return scyna.OK
}
