package model

import scyna "github.com/scyna/core"

var (
	MODULE_EXISTED        = scyna.NewError(100, "Module Existed")
	MODULE_NOT_EXIST      = scyna.NewError(101, "Module Not Exist")
	APPLICATION_EXISTED   = scyna.NewError(102, "Application Existed")
	APPLICATION_NOT_EXIST = scyna.NewError(103, "Application Not Exist")
	CLIENT_EXISTED        = scyna.NewError(104, "Client Existed")
	CLIENT_NOT_EXIST      = scyna.NewError(105, "Client Not Exist")
	SETTING_NOT_EXIST     = scyna.NewError(106, "Setting Not Exist")
	SETTING_EXISTED       = scyna.NewError(107, "Setting Existed")
)
