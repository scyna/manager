package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	scyna_const "github.com/scyna/core/const"
	"github.com/scyna/manager/model"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func CreateClientHandler(ctx *scyna.Endpoint, request *proto.CreateClientRequest) scyna.Error {
	if err := validation.ValidateStruct(request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Secret, validation.Required),
		validation.Field(&request.Name, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	if _, err := repository.GetClient(request.Id); err == nil {
		return model.CLIENT_EXISTED
	}

	client := repository.Client{
		ID:     request.Id,
		Name:   request.Name,
		Secret: request.Secret,
	}

	if err := repository.CreateClient(&client); err != nil {
		return err
	}

	scyna.EmitSignal(scyna_const.CLIENT_UPDATE_CHANNEL, &proto.UpdateClientSignal{
		Id: request.Id,
	})

	return scyna.OK
}
