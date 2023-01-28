package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/model"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func CreateModuleHandler(ctx *scyna.Endpoint, request *proto.CreateModuleRequest) scyna.Error {
	repository := repository.LoadRepository(ctx.Logger)

	if err := validateCreateModuleRequest(request); err != nil {
		return scyna.REQUEST_INVALID
	}

	if _, err := repository.GetModule(request.Code); err == nil {
		return model.MODULE_EXISTED
	}

	module := model.Module{
		Code:   request.Code,
		Secret: request.Secret,
	}

	if err := repository.CreateModule(&module); err != nil {
		return err
	}

	return scyna.OK
}

func validateCreateModuleRequest(request *proto.CreateModuleRequest) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Code, validation.Required, validation.Length(1, 50)),
		validation.Field(&request.Secret, validation.Required, validation.Length(8, 50)))
}
