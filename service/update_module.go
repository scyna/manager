package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func UpdateModuleHandler(ctx *scyna.Endpoint, request *proto.UpdateModuleRequest) scyna.Error {

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Code, validation.Required),
		validation.Field(&request.Name, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	if _, err := repository.GetModule(request.Code); err != nil {
		return err
	}

	module := repository.Module{
		Code: request.Code,
		Name: request.Name,
	}

	if err := repository.UpdateModule(&module); err != nil {
		return err
	}

	return scyna.OK
}
