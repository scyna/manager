package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func GetSettingHandler(ctx *scyna.Endpoint, request *proto.GetSettingRequest) scyna.Error {

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Module, validation.Required),
		validation.Field(&request.Key, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	setting, err := repository.GetSetting(request.Module, request.Key)
	if err != nil {
		return err
	}

	return ctx.OK(&proto.GetSettingResponse{
		Setting: &proto.Setting{
			Module: setting.Module,
			Key:    setting.Key,
			Value:  setting.Value,
		}})
}
