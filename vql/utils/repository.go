package utils

import (
	"errors"

	"github.com/Cyarun/CyFir/constants"
	"github.com/Cyarun/CyFir/services"
	"www.velocidex.com/golang/vfilter"
)

func GetRepository(scope vfilter.Scope) (services.Repository, error) {
	any_obj, pres := scope.Resolve(constants.SCOPE_REPOSITORY)
	if !pres {
		return nil, errors.New("Repository not found in scope!!")
	}
	repository, ok := any_obj.(services.Repository)
	if !ok {
		return nil, errors.New("Repository not found in scope!!")
	}
	return repository, nil
}
