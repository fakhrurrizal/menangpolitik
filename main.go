package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"polintar/config"
	"polintar/router"
	"polintar/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	// "github.com/gofiber/template/html"
)

var BasicData fiber.Map

func main() {

	env := config.LoadConfig().Environtment

	engine := html.NewFileSystem(http.Dir("./views"), ".html")

	engine.Reload(true) // Optional. Default: false
	if env == "development" {
		engine.Debug(true) // Optional. Default: false
	}
	engine.Layout("embed") // Optional. Default: "embed"

	engine.Delims("{{", "}}") // Optional. Default: engine delimiters

	app := fiber.New(fiber.Config{
		Views:             engine,
		PassLocalsToViews: true,
	})
	// if env == "production" {
	// 	app.Use(redirectHTTPToHTTPS())
	// }
	app.Use(func(c *fiber.Ctx) error {
		isDevelopment := false
		if env == "development" {
			isDevelopment = true
		}
		if env == "maintenance" {
			app.Static("assets", "/public")
			return c.Render("components/maintenance", fiber.Map{})
		}
		h := utils.Host(c)
		// fmt.Println("host :", h)
		c.Locals("layoutData", fiber.Map{
			"FrontEndUrl":   "https://menangpolitik.id",
			"ApiUrl":        "https://menangpolitik.id/api",
			"AppId":         config.LoadConfig().AppId,
			"Company":       h.CompanyID,
			"IsDevelopment": isDevelopment,
			"HostName":      c.Hostname(),
			"Logo":          "https://politikpintar.id/api/assets/uploads/logo/285/2023-08/1692326750.png",
			"CompanyEmail":  "menangpolitik@gmail.com",
			"CompanyPhone":  "089977665544",
		})
		return c.Next()
	})
	router.Init(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}

// func redirectHTTPToHTTPS() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		h := utils.Host(c)
// 		if c.Protocol() == "http" {
// 			// fmt.Println("host access:", c.Hostname())
// 			return c.Redirect(fmt.Sprintf("%s%s", h.FrontendURL, c.OriginalURL()))
// 		}
// 		return c.Next()
// 	}
// }
