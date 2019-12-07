package controller

import (
	"az-keyvault-secrets-operator/pkg/controller/keyvaultmap"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, keyvaultmap.Add)
}
