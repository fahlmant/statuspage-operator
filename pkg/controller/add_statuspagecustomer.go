package controller

import (
	"github.com/fahlmant/statuspage-operator/pkg/controller/statuspagecustomer"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, statuspagecustomer.Add)
}
