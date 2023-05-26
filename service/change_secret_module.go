package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func ChangeSecretModuleHandler(ctx *scyna.Endpoint, request *proto.ChangeSecretModuleRequest) scyna.Error {

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Code, validation.Required),
		validation.Field(&request.Secret, validation.Required),
		validation.Field(&request.Password, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	if _, err := repository.GetModule(request.Code); err != nil {
		return err
	}

	if err := repository.ChangePasswordModule(request.Code, request.Secret); err != nil {
		return err
	}

	return scyna.OK
}
