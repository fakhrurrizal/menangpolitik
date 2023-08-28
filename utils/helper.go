package utils

import (
	"os"
	"polintar/config"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
)

func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type Data fiber.Map

func BasicData() *Data {
	FronEndUrl := ""
	ApiUrl := ""
	return &Data{
		ApiUrl:     config.LoadConfig().FrontEndUrl,
		FronEndUrl: config.LoadConfig().BaseUrl,
	}
}

func TrimToMaxChars(input string, maxChars int) string {
	if utf8.RuneCountInString(input) > maxChars {
		return input[:maxChars]
	}
	return input
}
