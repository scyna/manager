package main

import (
	"flag"
	"github.com/scyna/manager/domain"
	"github.com/scyna/manager/repository"
	"github.com/scyna/manager/service"

	scyna "github.com/scyna/core"
)

const MODULE_CODE = "scyna.manager"

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

	domain.AttachRepositoryCreator(repository.NewRepository)
	scyna.RegisterEndpoint(service.CREATE_MODULE_URL, service.CreateModuleHandler)

	scyna.Start()
}
