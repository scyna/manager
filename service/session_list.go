package service

import (
	"strconv"
	"time"

	scyna "github.com/scyna/core"
	"github.com/scyna/manager/domain"
	"github.com/scyna/manager/model"
	proto "github.com/scyna/manager/proto/generated"
)

func ListSessionHandler(ctx *scyna.Endpoint, request *proto.ListSessionRequest) scyna.Error {
	service := domain.NewModuleService(&ctx.Context)

	var module *model.Module
	if ret, err := service.GetModule(request.ModuleCode); err != nil {
		return err
	} else {
		module = ret
	}

	var sessions []model.Session
	if ret, err := service.Repository.ListSession(module); err != nil {
		return err
	} else {
		sessions = ret
	}

	var response = proto.ListSessionResponse{}
	for _, session := range sessions {
		response.Sessions = append(response.Sessions, &proto.SessionItem{
			SessionID:   strconv.FormatInt(session.Id, 10),
			State:       int32(session.State),
			LastUpdated: session.LastUpdated.Format(time.RFC3339),
			StartTime:   session.StartTime.Format(time.RFC3339),
			EndTime:     session.EndTime.Format(time.RFC3339),
		})
	}
	ctx.Response(&response)

	return scyna.OK
}
