package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	proto "github.com/scyna/manager/proto/generated"
)

func LoginHandler(ctx *scyna.Endpoint, request *proto.LoginRequest) scyna.Error {

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Password, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	token, expired, err := ctx.Authenticate("scyna_manager", []string{"scyna_manager"})
	if err != nil {
		return scyna.SERVER_ERROR
	}

	return ctx.AuthOK(token, expired, &proto.LoginResponse{Session: token})
}
