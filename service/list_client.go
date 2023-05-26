package service

import (
	scyna "github.com/scyna/core"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func ListClientHandler(ctx *scyna.Endpoint, request *proto.ListClientRequest) scyna.Error {

	clients, err := repository.ListClient()
	if err != nil {
		return err
	}

	var protoClients []*proto.Client
	for _, client := range clients {
		protoClients = append(protoClients, &proto.Client{
			Id:   client.ID,
			Name: client.Name,
		})
	}

	return ctx.OK(&proto.ListClientResponse{
		Clients: protoClients,
	})
}
