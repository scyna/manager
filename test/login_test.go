package test

import (
	"testing"

	scyna_test "github.com/scyna/core/testing"
	proto "github.com/scyna/manager/proto/generated"
	"github.com/scyna/manager/service"
)

func TestLogin_Success(t *testing.T) {
	scyna_test.Endpoint(service.LOGIN_EP).
		WithRequest(&proto.LoginRequest{
			Username: "test",
			Password: "123456",
		}).
		ExpectResponseLike(&proto.LoginResponse{}).Run(t)
}
