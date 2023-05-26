package test

import (
	"testing"

	scyna_test "github.com/scyna/core/testing"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/service"
)

func TestChangeSecretModule_Success(t *testing.T) {
	scyna_test.Endpoint(service.CHANGE_SECRET_MODULE_EP).
		WithRequest(&proto.ChangeSecretModuleRequest{
			Code: "test",
		}).
		ExpectSuccess().Run(t)
}
