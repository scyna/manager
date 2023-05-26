package main

import (
	"flag"

	scyna "github.com/scyna/core"
	"github.com/scyna/manager/service"
)

const MODULE_CODE = "scyna_manager"

func main() {
	secret_ := flag.String("secret", "123456", "Secret")
	managerUrl := flag.String("managerUrl", "http://127.0.0.1:8081", "Manager Url")
	flag.Parse()

	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: *managerUrl,
		Name:       MODULE_CODE,
		Secret:     *secret_,
	})
	scyna.UseRemoteLog(1)
	defer scyna.Release()

	scyna.RegisterEndpoint(service.LOGIN_EP, service.LoginHandler)
	scyna.RegisterEndpoint(service.CREATE_MODULE_EP, service.CreateModuleHandler)
	scyna.RegisterEndpoint(service.UPDATE_MODULE_EP, service.UpdateModuleHandler)
	scyna.RegisterEndpoint(service.CHANGE_SECRET_MODULE_EP, service.ChangeSecretModuleHandler)
	scyna.RegisterEndpoint(service.DELETE_MODULE_EP, service.DeleteModuleHandler)
	scyna.RegisterEndpoint(service.LIST_MODULE_EP, service.ListModuleHandler)
	scyna.RegisterEndpoint(service.GET_MODULE_EP, service.GetModuleHandler)
	scyna.RegisterEndpoint(service.UPDATE_CONFIG_MODULE_EP, service.UpdateConfigModuleHandler)
	scyna.RegisterEndpoint(service.CREATE_APPLICATION_EP, service.CreateApplicationHandler)
	scyna.RegisterEndpoint(service.UPDATE_APPLICATION_EP, service.UpdateApplicationHandler)
	scyna.RegisterEndpoint(service.DELETE_APPLICATION_EP, service.DeleteApplicationHandler)
	scyna.RegisterEndpoint(service.GET_APPLICATION_EP, service.GetApplicationHandler)
	scyna.RegisterEndpoint(service.LIST_APPLICATION_EP, service.ListApplicationHandler)
	scyna.RegisterEndpoint(service.CREATE_CLIENT_EP, service.CreateClientHandler)
	scyna.RegisterEndpoint(service.UPDATE_CLIENT_EP, service.UpdateClientHandler)
	scyna.RegisterEndpoint(service.DELETE_CLIENT_EP, service.DeleteClientHandler)
	scyna.RegisterEndpoint(service.GET_CLIENT_EP, service.GetClientHandler)
	scyna.RegisterEndpoint(service.LIST_CLIENT_EP, service.ListClientHandler)
	scyna.RegisterEndpoint(service.CHANGE_SECRET_CLIENT_EP, service.ChangeSecretClientHandler)
	scyna.Start()
}
