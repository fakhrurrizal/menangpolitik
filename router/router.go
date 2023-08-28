package router

import (
	"fmt"
	"os"
	"os/exec"
	"polintar/controllers"
	"polintar/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func Init(app *fiber.App) {
	utils.BasicData()
	app.Get("/", controllers.Home)
	app.Get("/wawasan", controllers.Wawasan)
	app.Get("/kecerdasan", controllers.Kecerdasan)
	app.Get("/pricing", controllers.Pricing)
	app.Get("/campaign", controllers.Kampanye)
	app.Get("/addon", controllers.AddOn)
	app.Get("/payment", controllers.Payment)
	app.Get("/login", controllers.Login)
	app.Get("/forgot-password", controllers.ForgotPassword)
	app.Get("/register", controllers.Register)
	app.Get("/invoice/:id", controllers.Invoice)
	app.Get("/transaction", controllers.Transaction)
	app.Get("/build", controllers.Build)
	app.Get("/version", controllers.Version)
	app.Get("/send-email", controllers.SendEmail)
	app.Get("/reset-password", controllers.ResetPassword)
	app.Get("/news-category/:id", controllers.NewCategory)
	app.Get("/news-tags/:id", controllers.NewsTags)
	app.Get("/signout-all", controllers.SignoutAll)
	app.Get("/privacy-policy", controllers.PrivacyPolicy)
	app.Get("/terms-condition", controllers.TermsCondition)

	// app.Get("/signin", controllers.Signin)

	app.Get("/verify-email", controllers.VerifyEmail)

	app.Static("/assets", "./dist/assets")
	app.Static("/images", "./dist/images")
	app.Static("/assets", "./public")

	app.All("/api/*", proxy.Balancer(proxy.Config{
		ModifyRequest: func(req *fiber.Ctx) error {
			req.Path(strings.TrimPrefix(req.Path(), "/api"))
			return nil
		},
		Servers: []string{os.Getenv("PROXY_API")},
	}))

	reactDevUrl := "http://localhost:5173/"
	if CheckUrlExist(reactDevUrl) {
		//menjalankan react development server di /app jika ada yarn dev
		reactHandler := proxy.Balancer(proxy.Config{
			ModifyRequest: func(req *fiber.Ctx) error {
				req.Path(strings.TrimPrefix(req.Path(), "/app"))
				return nil
			},
			Servers: []string{
				reactDevUrl,
			},
		})
		fmt.Println("proxy reactDev running")
		app.Get("/app/*", reactHandler)
		app.Get("/src/*", reactHandler)
		app.Get("*", reactHandler)
	}

	app.Static("/app", "./dist", fiber.Static{
		Index: "index.html",
	})
	app.Static("/builder", "./builder/layoutit", fiber.Static{
		Index: "index.html",
	})
	app.Get("/news/:category/:slug", controllers.NewDetail)
	app.Get("/app/*", func(c *fiber.Ctx) error {
		return c.SendFile("./dist/index.html")
	})

}

func CheckUrlExist(url string) (status bool) {
	cmd := exec.Command("curl", "-s", "-o", "/dev/null", "-w", "%{http_code}", url)
	output, err := cmd.Output()
	if err != nil {
		status = false
	}

	statusCode := strings.TrimSpace(string(output))
	if statusCode == "200" {
		status = true
	}

	fmt.Println(url, status)
	return
}
