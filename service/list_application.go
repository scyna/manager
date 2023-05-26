package service

import (
	scyna "github.com/scyna/core"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func ListApplicationHandler(ctx *scyna.Endpoint, request *proto.ListApplicationRequest) scyna.Error {

	applications, err := repository.ListApplication()
	if err != nil {
		return err
	}
	var response proto.ListApplicationResponse
	for _, m := range applications {
		response.Applications = append(response.Applications, &proto.Application{
			Code:     m.Code,
			Name:     m.Name,
			AuthPath: m.AuthPath,
		})
	}
	return ctx.OK(&response)
}
