package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func GetModuleHandler(ctx *scyna.Endpoint, request *proto.GetModuleRequest) scyna.Error {

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Code, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	module, err := repository.GetModule(request.Code)
	if err != nil {
		return err
	}

	config, err := repository.GetConfigModule(request.Code)
	if err != nil {
		return err
	}

	return ctx.OK(&proto.GetModuleResponse{
		Module: &proto.Module{
			Code: module.Code,
			Name: module.Name,
			Config: &proto.Config{
				NatsUrl:      config.NatsURL,
				NatsUsername: config.NatsUsername,
				NatsPassword: config.NatsPassword,
				DbHost:       config.DBHost,
				DbUsername:   config.DBUsername,
				DbPassword:   config.DBPassword,
				DbLocation:   config.DBLocation,
			},
		},
	})

}
