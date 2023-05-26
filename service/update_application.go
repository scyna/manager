package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	scyna_const "github.com/scyna/core/const"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func UpdateApplicationHandler(ctx *scyna.Endpoint, request *proto.UpdateApplicationRequest) scyna.Error {

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Code, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.AuthPath, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	app, err := repository.GetApplication(request.Code)
	if err != nil {
		return err
	}

	app.Name = request.Name
	app.AuthPath = request.AuthPath

	if err := repository.UpdateApplication(app); err != nil {
		return err
	}

	scyna.EmitSignal(scyna_const.APP_UPDATE_CHANNEL, &proto.UpdateApplicationSignal{Code: app.Code})

	return scyna.OK
}
