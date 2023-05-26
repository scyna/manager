package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	scyna_const "github.com/scyna/core/const"
	scyna_proto "github.com/scyna/core/proto/generated"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func DeleteSettingHandler(ctx *scyna.Endpoint, request *proto.DeleteSettingRequest) scyna.Error {

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Key, validation.Required),
		validation.Field(&request.Module, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	if _, err := repository.GetSetting(request.Module, request.Key); err != nil {
		return err
	}

	if err := repository.DeleteSetting(request.Module, request.Key); err != nil {
		return err
	}

	scyna.EmitSignal(scyna_const.SETTING_UPDATE_CHANNEL, &scyna_proto.SettingUpdatedSignal{
		Module: request.Module,
		Key:    request.Key,
		Value:  "",
	})

	return scyna.OK
}
