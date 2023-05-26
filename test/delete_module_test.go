package test

import (
	"testing"

	scyna_test "github.com/scyna/core/testing"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/service"
)

func TestDeleteModule_Success(t *testing.T) {
	scyna_test.Endpoint(service.DELETE_MODULE_EP).
		WithRequest(&proto.DeleteModuleRequest{
			Code:     "test",
			Password: "123456",
		}).
		ExpectSuccess().Run(t)
}
