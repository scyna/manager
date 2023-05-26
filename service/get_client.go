package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	scyna "github.com/scyna/core"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func GetClientHandler(ctx *scyna.Endpoint, request *proto.GetClientRequest) scyna.Error {
	if err := validation.ValidateStruct(request,
		validation.Field(&request.Id, validation.Required),
	); err != nil {
		ctx.Error(err.Error())
		return scyna.REQUEST_INVALID
	}

	client, err := repository.GetClient(request.Id)
	if err != nil {
		return err
	}

	return ctx.OK(&proto.GetClientResponse{
		Client: &proto.Client{
			Id:   client.ID,
			Name: client.Name},
	})
}
