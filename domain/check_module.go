package domain

import (
	scyna "github.com/scyna/core"
	"github.com/scyna/manager/model"
)

func AssureModuleNotExists(repository IRepository, code string) scyna.Error {
	if _, err := repository.GetModule(code); err == nil {
		return model.MODULE_EXISTED
	}
	return nil
}
