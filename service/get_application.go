package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func GetApplicationHandler(ctx *scyna.Endpoint, request *proto.GetApplicationRequest) scyna.Error {

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Code, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	app, err := repository.GetApplication(request.Code)
	if err != nil {
		return err
	}

	ctx.Response(&proto.GetApplicationResponse{
		Application: &proto.Application{
			Code:     app.Code,
			Name:     app.Name,
			AuthPath: app.AuthPath,
		},
	})

	return scyna.OK
}
