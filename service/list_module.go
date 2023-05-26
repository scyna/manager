package service

import (
	scyna "github.com/scyna/core"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/repository"
)

func ListModuleHandler(ctx *scyna.Endpoint, request *proto.ListModuleRequest) scyna.Error {

	modules, err := repository.ListModule()
	if err != nil {
		return err
	}
	var response proto.ListModuleResponse
	for _, m := range modules {
		response.Modules = append(response.Modules, &proto.Module{
			Code: m.Code,
			Name: m.Name,
		})
	}
	return ctx.OK(&response)
}
