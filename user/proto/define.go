package proto

import (
	scyna "github.com/scyna/core"
)

const (
	CREATE_USER_URL      = "/scyman/user/create"
	GET_USER_URL         = "/scyman/user/get"
	USER_CREATED_CHANNEL = "scyman.user.user_created"
)

var (
	USER_EXISTED     = &scyna.Error{Code: 100, Message: "User Existed"}
	USER_NOT_EXISTED = &scyna.Error{Code: 101, Message: "User Not Existed"}
)
