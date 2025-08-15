package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yayayapluto/api_lelang_online_golang_revisi_2/internal/api/routes"
	"github.com/yayayapluto/api_lelang_online_golang_revisi_2/internal/utils"
	"gorm.io/gorm"
)

func NewApp(db *gorm.DB) (*fiber.App, error) {
	utils.InitValidator()
	//validator := utils.Validate

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	routeConfig := routes.RouteConfig{
		App: app,
	}
	routeConfig.Setup()

	return app, nil
}
