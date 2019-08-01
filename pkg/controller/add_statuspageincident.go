package controller

import (
	"github.com/fahlmant/statuspage-operator/pkg/controller/statuspageincident"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, statuspageincident.Add)
}
