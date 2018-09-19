package models

import "app/yast/services"

func BootstrapModels() {
	services.Logger.Info("Bootstrapping Models")
	services.DB.AutoMigrate(&User{})
	services.DB.AutoMigrate(&Profile{})
}
