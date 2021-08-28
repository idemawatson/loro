package controllers

import (
	"loro-handler/database"
	"loro-handler/repository"
)

var sharedInstance Controllers = *newControllers()

type Controllers struct {
	UserController UserController
}

func newControllers() *Controllers {
	mainTable := database.GetMainTable()
	uc := NewUserController(repository.NewUserRepository(&mainTable))
	return &Controllers{UserController: *uc}
}

func GetControllers() *Controllers {
	return &sharedInstance
}
