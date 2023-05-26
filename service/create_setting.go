package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	scyna_const "github.com/scyna/core/const"
	scyna_proto "github.com/scyna/core/proto/generated"
	"github.com/scyna/manager/model"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func CreateSettingHandler(ctx *scyna.Endpoint, request *proto.CreateSettingRequest) scyna.Error {

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Module, validation.Required),
		validation.Field(&request.Key, validation.Required),
		validation.Field(&request.Value, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	if _, err := repository.GetModule(request.Module); err != nil {
		return model.MODULE_NOT_EXIST
	}

	if _, err := repository.GetSetting(request.Module, request.Key); err == nil {
		return model.SETTING_EXISTED
	}

	if err := repository.CreateSetting(&repository.Setting{
		Module: request.Module,
		Key:    request.Key,
		Value:  request.Value,
	}); err != nil {
		return err
	}

	scyna.EmitSignal(scyna_const.SETTING_UPDATE_CHANNEL, &scyna_proto.SettingUpdatedSignal{
		Module: request.Module,
		Key:    request.Key,
		Value:  request.Value,
	})

	return scyna.OK
}
