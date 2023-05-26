package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	scyna_const "github.com/scyna/core/const"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func DeleteClientHandler(ctx *scyna.Endpoint, request *proto.DeleteClientRequest) scyna.Error {

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Id, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	if _, err := repository.GetClient(request.Id); err != nil {
		return err
	}

	if err := repository.DeleteClient(request.Id); err != nil {
		return err
	}

	scyna.EmitSignal(scyna_const.CLIENT_UPDATE_CHANNEL, &proto.UpdateClientSignal{
		Id: request.Id,
	})

	return scyna.OK
}
