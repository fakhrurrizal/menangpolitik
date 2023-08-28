package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"polintar/config"

	"github.com/gofiber/fiber/v2"
)

func VerifyEmail(c *fiber.Ctx) error {
	token := c.Query("token")
	url := config.LoadConfig().BaseUrl + "/v1/auth/verify-email"
	data := map[string]interface{}{
		"token": token,
	}
	fmt.Println(data)
	jsonData, _ := json.Marshal(data)

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		return c.Redirect(config.LoadConfig().FrontEndUrl+"/login?verify=success", fiber.StatusFound)
	}
	if response.StatusCode == 422 || response.StatusCode == 500 {
		body, _ := ioutil.ReadAll(response.Body)
		return c.Status(response.StatusCode).JSON(body)
	}
	return c.Redirect(config.LoadConfig().FrontEndUrl+"/resend_email", fiber.StatusFound)
}
