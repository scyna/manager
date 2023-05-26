package service

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/model"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func CreateModuleHandler(ctx *scyna.Endpoint, request *proto.CreateModuleRequest) scyna.Error {

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Code, validation.Required),
		validation.Field(&request.Secret, validation.Required),
		validation.Field(&request.Name, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	if _, err := repository.GetModule(request.Code); err == nil {
		return model.MODULE_EXISTED
	}

	module := repository.Module{
		Code:   request.Code,
		Name:   request.Name,
		Secret: request.Secret,
	}

	config := repository.Config{
		NatsURL:      request.Config.NatsUrl,
		NatsUsername: request.Config.NatsUsername,
		NatsPassword: request.Config.NatsPassword,
		DBHost:       request.Config.DbHost,
		DBUsername:   request.Config.DbUsername,
		DBPassword:   request.Config.DbPassword,
		DBLocation:   request.Config.DbLocation,
	}

	val, err := json.Marshal(config)
	if err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	if err := repository.CreateModule(&module, string(val)); err != nil {
		return err
	}

	return scyna.OK
}
