/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package api

import (
	"fmt"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/carbonaut/docs"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Port string
}

// @title Carbonaut API
// @version 0.1
// @description This API is used to interact with Carbonaut resources
// @license.name MIT
// @license.url https://mit-license.org
func Start(c *Config) error {
	app := fiber.New()
	version := "v0.1"
	v := app.Group(fmt.Sprintf("/api/%s", version))
	docs.SwaggerInfo.Version = version

	// Base path handlers
	addBasePathRoutes(v)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	log.Info().Msgf("Swagger endpoint: http://127.0.0.1:%s/api/%s/swagger/index.html", c.Port, version)
	if err := app.Listen(fmt.Sprintf(":%s", c.Port)); err != nil {
		return err
	}
	return nil
}

func addBasePathRoutes(r fiber.Router) {
	r.Get("/status", statusHandler)
	r.Post("/init", initHandler)
	// host a add swagger web UI
	r.Get("/swagger/*", swagger.HandlerDefault)
}

const statusOK = "Carbonaut API is running!"

// @description Carbonaut Status Endpoint
// @Success 200 {string} Carbonaut API is running!
// @Router /api/v0.1/status/ [get]
func statusHandler(c *fiber.Ctx) error {
	return c.SendString(statusOK)
}

// @description Initialize carbonaut to be fully functional
// @Success 200 {string} Carbonaut initialized!
// @Router /api/v0.1/init [post]
func initHandler(c *fiber.Ctx) error {
	return c.SendString("wip, not implemented")
}
