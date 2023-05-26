package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	scyna_const "github.com/scyna/core/const"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func DeleteApplicationHandler(ctx *scyna.Endpoint, request *proto.DeleteApplicationRequest) scyna.Error {

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Code, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	_, err := repository.GetApplication(request.Code)
	if err != nil {
		return err
	}

	if err := repository.DeleteApplication(request.Code); err != nil {
		return err
	}

	scyna.EmitSignal(scyna_const.APP_UPDATE_CHANNEL, &proto.UpdateApplicationSignal{Code: request.Code})

	return scyna.OK
}
