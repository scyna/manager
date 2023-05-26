package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/model"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func ListSettingHandler(ctx *scyna.Endpoint, request *proto.ListSettingRequest) scyna.Error {

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Module, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	if _, err := repository.GetModule(request.Module); err != nil {
		return model.MODULE_NOT_EXIST
	}

	settings, err := repository.ListSetting(request.Module)
	if err != nil {
		return err
	}
	var response proto.ListSettingResponse
	for _, m := range settings {
		response.Settings = append(response.Settings, &proto.Setting{
			Module: m.Module,
			Key:    m.Key,
			Value:  m.Value,
		})
	}
	return ctx.OK(&response)
}
