package test

import (
	"testing"

	scyna_test "github.com/scyna/core/testing"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/service"
)

func TestUpdateModule_Success(t *testing.T) {
	scyna_test.Endpoint(service.UPDATE_MODULE_EP).
		WithRequest(&proto.UpdateModuleRequest{
			Code: "test",
			Name: "test",
		}).
		ExpectSuccess().Run(t)
}
