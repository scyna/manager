package model

import scyna "github.com/scyna/core"

var (
	MODULE_EXISTED   = scyna.NewError(100, "Module Existed")
	MODULE_NOT_EXIST = scyna.NewError(101, "Module Not Exist")
)
