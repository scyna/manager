package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	scyna_const "github.com/scyna/core/const"
	"github.com/scyna/manager/model"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func CreateApplicationHandler(ctx *scyna.Endpoint, request *proto.CreateApplicationRequest) scyna.Error {

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Code, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.AuthPath, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	if _, err := repository.GetApplication(request.Code); err == nil {
		return model.APPLICATION_EXISTED
	}

	app := repository.Application{
		Code:     request.Code,
		Name:     request.Name,
		AuthPath: request.AuthPath,
	}

	if err := repository.CreateApplication(&app); err != nil {
		return err
	}

	scyna.EmitSignal(scyna_const.APP_UPDATE_CHANNEL, &proto.UpdateApplicationSignal{Code: app.Code})

	return scyna.OK
}
