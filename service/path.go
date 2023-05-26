package service

const BASEPATH = "/scyna2"

const (
	CREATE_DOMAIN_EP        = BASEPATH + "/manager/domain/create"
	CREATE_MODULE_EP        = BASEPATH + "/manager/module/create"
	UPDATE_MODULE_EP        = BASEPATH + "/manager/module/update"
	CHANGE_SECRET_MODULE_EP = BASEPATH + "/manager/module/change-secret"
	DELETE_MODULE_EP        = BASEPATH + "/manager/module/delete"
	LIST_MODULE_EP          = BASEPATH + "/manager/module/list"
	GET_MODULE_EP           = BASEPATH + "/manager/module/get"
	UPDATE_CONFIG_MODULE_EP = BASEPATH + "/manager/module/update-config"
	LOGIN_EP                = BASEPATH + "/manager/login"

	CREATE_APPLICATION_EP = BASEPATH + "/manager/application/create"
	UPDATE_APPLICATION_EP = BASEPATH + "/manager/application/update"
	DELETE_APPLICATION_EP = BASEPATH + "/manager/application/delete"
	LIST_APPLICATION_EP   = BASEPATH + "/manager/application/list"
	GET_APPLICATION_EP    = BASEPATH + "/manager/application/get"

	CREATE_CLIENT_EP        = BASEPATH + "/manager/client/create"
	UPDATE_CLIENT_EP        = BASEPATH + "/manager/client/update"
	DELETE_CLIENT_EP        = BASEPATH + "/manager/client/delete"
	LIST_CLIENT_EP          = BASEPATH + "/manager/client/list"
	GET_CLIENT_EP           = BASEPATH + "/manager/client/get"
	CHANGE_SECRET_CLIENT_EP = BASEPATH + "/manager/client/change-secret"
)
