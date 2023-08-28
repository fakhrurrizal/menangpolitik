package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"polintar/config"
	"polintar/entity"
	"polintar/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	h := utils.Host(c)

	view := h.ViewFolder + "/pages/index"
	return c.Render(view, fiber.Map{
		"SiteUrlPage":  h.FrontEndUrl + "/",
		"CurrentRoute": "/",
		"Host":         h.FrontEndUrl,
	}, Layout(c))
}

func Wawasan(c *fiber.Ctx) error {

	host := c.Hostname()
	return c.Render(Pages(c)+"wawasan", fiber.Map{
		"SiteUrlPage":  host + "/",
		"CurrentRoute": "/wawasan",
		"Title":        "Wawasan | Menangpol",
	}, Layout(c))
}

func Signin(c *fiber.Ctx) error {
	h := utils.Host(c)
	view := h.ViewFolder + "/../src/pages/signin"
	return c.Render(view, fiber.Map{
		"SiteUrlPage":  h.FrontEndUrl + "/signin",
		"CurrentRoute": "/signin",
		"Title":        Login,
		"Host":         h.FrontEndUrl,
	})
}

func Kecerdasan(c *fiber.Ctx) error {

	return c.Render(Pages(c)+"kecerdasan", fiber.Map{
		"CurrentRoute": "/kecerdasan",
		"Title":        "Produk | Menangpol",
		"SiteUrlPage":  "https://politikpintar.id/pricing",
	}, Layout(c))
}

func Pricing(c *fiber.Ctx) error {

	return c.Render(Pages(c)+"pricing", fiber.Map{
		"CurrentRoute": "/pricing",
		"Description":  "List harga politik pintar",
		"SiteUrlPage":  "https://politikpintar.id/pricing",

		"Title": "Harga | Menangpol",
	}, Layout(c))
}

func Kampanye(c *fiber.Ctx) error {

	return c.Render(Pages(c)+"kampanye", fiber.Map{
		"CurrentRoute": "/campaign",
		"Description":  "List harga politik pintar",
		"SiteUrlPage":  "https://politikpintar.id/campaign",

		"Title": "Kampanye | Menangpol",
	}, Layout(c))
}

func AddOn(c *fiber.Ctx) error {

	return c.Render(Pages(c)+"addon", fiber.Map{
		"CurrentRoute": "/addon",
		"Description":  "List addon politik pintar",
		"SiteUrlPage":  "https://politikpintar.id/addon",

		"Title": "Addon Paket | Menangpol",
	}, Layout(c))
}

func Payment(c *fiber.Ctx) error {

	return c.Render(Pages(c)+"payment", fiber.Map{
		"CurrentRoute": "/payment",
		"SiteUrlPage":  "https://politikpintar.id/payment",

		"Title": "Pembayaran Paket | Menangpol",
	}, Layout(c))
}

func Invoice(c *fiber.Ctx) error {
	id := c.Params("id")

	return c.Render(Pages(c)+"invoice", fiber.Map{
		"CurrentRoute": "/invoice",
		"SiteUrlPage":  "https://politikpintar.id/payment",

		"Title":      "Invoice " + id + " | Menangpol",
		"Invoice_id": id,
	}, Layout(c))
}

func Transaction(c *fiber.Ctx) error {

	return c.Render(Pages(c)+"transaction", fiber.Map{
		"CurrentRoute": "/transaction",
		"SiteUrlPage":  "https://politikpintar.id/transaction",

		"Title": "Transaction | Menangpol",
	}, Layout(c))
}

func Login(c *fiber.Ctx) error {

	next := c.Query("next")
	target := "/"

	if next != "" {
		target = next
	}
	return c.Render(Pages(c)+"login", fiber.Map{
		"CurrentRoute": "/login",
		"NextLocation": target,

		"Title": "Login | Menangpol",
	}, Layout(c))
}

func Register(c *fiber.Ctx) error {

	return c.Render(Pages(c)+"register", fiber.Map{
		"CurrentRoute": "/register",

		"Title": "Register | Menangpol",
	}, Layout(c))
}

func ForgotPassword(c *fiber.Ctx) error {

	return c.Render(Pages(c)+"forgot-password", fiber.Map{
		"CurrentRoute": "/forgot-password",
		"NextLocation": "/send-email",

		"Title": "Forgot Password | Menangpol",
	})
}

func PrivacyPolicy(c *fiber.Ctx) error {

	return c.Render(Pages(c)+"privacy-policy", fiber.Map{
		"CurrentRoute": "/privacy-policy",
		"Description":  "Kebijakan Privasi",
		// "SiteUrlPage":  "https://politikpintar.id/pricing",

		"Title": "Kebijakan Privasi | Menangpol",
	}, Layout(c))
}

func TermsCondition(c *fiber.Ctx) error {

	return c.Render(Pages(c)+"terms-condition", fiber.Map{
		"CurrentRoute": "/terms-condition",
		"Description":  "Kebijakan Privasi",
		// "SiteUrlPage":  "https://politikpintar.id/pricing",

		"Title": "Syarat & Ketentuan | Menangpol",
	}, Layout(c))
}

func SendEmail(c *fiber.Ctx) error {

	return c.Render(Pages(c)+"send-email", fiber.Map{
		"CurrentRoute": "/send-email",

		"Title": "Send email link | Menangpol",
	})
}

func VerifiedEmail(c *fiber.Ctx) error {

	return c.Render(Pages(c)+"verify-email", fiber.Map{
		"CurrentRoute": "/verify-email",

		"Title": "Send email Verify | Menangpol",
	})
}

func ResetPassword(c *fiber.Ctx) error {

	return c.Render(Pages(c)+"reset-password", fiber.Map{
		"CurrentRoute": "/reset-password",

		"Title": "Reset Password | Menangpol",
	})
}

func NewDetail(c *fiber.Ctx) error {
	slug := c.Params("slug")
	category := c.Params("category")
	if slug == "" && category != "" {
		slug = category
	}
	url := config.LoadConfig().BaseUrl + "/v1/landing-page/news/" + slug
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return c.JSON(fiber.Map{
			"status":  404,
			"details": "News Not Found",
		})
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return c.JSON(fiber.Map{
			"status":  500,
			"details": "Body News InValid",
		})
	}
	var data entity.NewsDetail

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return c.JSON(fiber.Map{
			"status":  500,
			"details": "Body News InValid",
		})
	}
	createdAt := data.Data.CreatedAt
	datetime, err := time.Parse(time.RFC3339Nano, createdAt)
	if err == nil {
		createdAt = datetime.Format("02 January 2006 03:04 PM")
	}

	description := data.Data.Description
	description = strings.ReplaceAll(description, "\n", "")
	description = strings.ReplaceAll(description, "\r", "")
	description = strings.ReplaceAll(description, `"`, "")
	description = utils.TrimToMaxChars(description, 155)

	tagsString := strings.Join(data.Data.Tags, ", ")

	CategoryNews := data.Data.Category.Label

	h := utils.Host(c)
	return c.Render(Pages(c)+"news-detail", fiber.Map{
		"ID":           data.Data.ID,
		"CurrentRoute": "/news/" + CategoryNews + slug,
		"Description":  description,
		"Body":         template.HTML(data.Data.Body),
		"SiteUrlPage":  h.FrontEndUrl + "/news/" + CategoryNews + slug,

		"Keyword":      data.Data.Keyword,
		"AuthorSite":   h.FrontEndUrl,
		"Title":        data.Data.Title,
		"Thumbnail":    data.Data.Image,
		"Tags":         tagsString,
		"CategoryName": data.Data.Category.Label,
		"CategoryId":   data.Data.Category.ID,
		"PublishAt":    data.Data.PublishAt,
		"ModifiedAt":   data.Data.UpdatedAt,
		"CreatedAt":    createdAt,
		"Author":       data.Data.Author,
		"Shared":       data.Data.Shared,
	}, Layout(c))
}

func NewCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	return c.Render(Pages(c)+"news-category", fiber.Map{
		"CurrentRoute": "/category/" + id,

		"Title":       "Category ",
		"Category_id": id,
	}, Layout(c))
}

func NewsTags(c *fiber.Ctx) error {
	id := c.Params("id")

	return c.Render(Pages(c)+"news-tags", fiber.Map{
		"CurrentRoute": "/tags/" + id,

		"Title":    "tags " + id,
		"news_tag": id,
	}, Layout(c))
}

func SignoutAll(c *fiber.Ctx) error {

	return c.Render(Pages(c)+"signout-all", fiber.Map{
		"CurrentRoute": "/singout-all",

		"Title": "Signout All | Menangpol",
	})
}
