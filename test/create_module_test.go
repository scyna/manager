package test

import (
	"testing"

	scyna_test "github.com/scyna/core/testing"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/service"
)

func TestCreateModule(t *testing.T) {
	scyna_test.Endpoint(service.CREATE_MODULE_EP).
		WithRequest(&proto.CreateModuleRequest{
			Code:   "scyna_test",
			Name:   "Scyna Test",
			Secret: "123456",
		}).
		ExpectSuccess().Run(t)
}
