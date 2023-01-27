package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/domain"
	"github.com/scyna/manager/model"
	proto "github.com/scyna/manager/proto/generated"
)

func CreateModuleHandler(ctx *scyna.Endpoint, request *proto.CreateModuleRequest) scyna.Error {

	repository := domain.LoadRepository(ctx.Logger)

	if err := validateCreateModuleRequest(request); err != nil {
		return scyna.REQUEST_INVALID
	}

	var ret scyna.Error
	if ret = domain.AssureModuleNotExists(repository, request.Code); ret != nil {
		return ret
	}

	module := model.Module{
		Code: request.Code,
	}

	if module.Secret, ret = model.ParseSecret(request.Secret); ret != nil {
		ctx.Logger.Error("wrong password")
		return ret
	}

	if ret = repository.CreateModule(ctx.Logger, &module); ret != nil {
		return ret
	}

	return scyna.OK
}

func validateCreateModuleRequest(request *proto.CreateModuleRequest) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Code, validation.Required, validation.Length(1, 50)),
		validation.Field(&request.Secret, validation.Required, validation.Length(8, 50)))
}
