package test

import (
	"os"
	"testing"

	scyna "github.com/scyna/core"
	scyna_test "github.com/scyna/core/testing"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/service"
)

var module *proto.Module
var modules []*proto.Module

func TestMain(m *testing.M) {
	scyna_test.Init("scyna_test")

	scyna.RegisterEndpoint(service.LOGIN_EP, service.LoginHandler)
	scyna.RegisterEndpoint(service.CREATE_MODULE_EP, service.CreateModuleHandler)
	scyna.RegisterEndpoint(service.UPDATE_MODULE_EP, service.UpdateModuleHandler)
	scyna.RegisterEndpoint(service.CHANGE_SECRET_MODULE_EP, service.ChangeSecretModuleHandler)
	scyna.RegisterEndpoint(service.DELETE_MODULE_EP, service.DeleteModuleHandler)
	scyna.RegisterEndpoint(service.LIST_MODULE_EP, service.ListModuleHandler)
	scyna.RegisterEndpoint(service.GET_MODULE_EP, service.GetModuleHandler)

	createModule()

	exitVal := m.Run()
	cleanup()
	scyna_test.Release()
	os.Exit(exitVal)
}

func createModule() {
	module = &proto.Module{
		Code: "test",
		Name: "test",
	}
	modules = []*proto.Module{
		{Code: "test1", Name: "test1"},
		{Code: "test2", Name: "test2"},
		{Code: "test3", Name: "test3"},
	}
}

func cleanup() {
	scyna.DB.Query("TRUNCATE scyna2.module", nil).ExecRelease()
}
