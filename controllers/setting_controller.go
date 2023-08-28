package controllers

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"os/exec"
	"polintar/config"
	"polintar/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CommandResult struct {
	Command string `json:"command"`
	Output  string `json:"output"`
	Error   string `json:"error"`
}

func Layout(c *fiber.Ctx) string {
	h := utils.Host(c)
	return h.ViewFolder + "/layouts/main"
}

func Pages(c *fiber.Ctx) string {
	h := utils.Host(c)
	return h.ViewFolder + "/pages/"
}

func Build(c *fiber.Ctx) error {
	origin := c.Query("origin")
	yarn := c.Query("yarn")
	if origin == "" {
		origin = "master"
	}
	commands := []string{}

	commands = append(commands, "git reset --hard")
	commands = append(commands, "git pull origin "+origin)
	email := false
	if yarn == "true" {
		commands = append(commands, "yarn install")
		commands = append(commands, "yarn build")
		email = true
	}
	fmt.Println(commands)
	var results []CommandResult

	if !email {
		results = executeCommand(commands, false)
		return c.Status(200).JSON(map[string]interface{}{
			"data": map[string]interface{}{
				"output": results,
				"origin": origin,
			},
			"message": "",
		})
	} else {
		go executeCommand(commands, true)
		return c.Status(200).JSON(map[string]interface{}{
			"message": "Result Will Be Emailed",
		})
	}
}

func Version(c *fiber.Ctx) error {
	commitID, timestamp, err := getLastCommitInfo()
	if err != nil {
		fmt.Println("Error:", err)
		return c.Status(500).JSON(map[string]interface{}{
			"message": "failed get version build",
		})
	}
	formattedTimestamp, err := convertTimestamp(timestamp)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return c.JSON(map[string]interface{}{
		"data": map[string]interface{}{
			"build_id":   commitID,
			"last_build": timestamp,
			"version":    formattedTimestamp,
			"host":       utils.Host(c),
			"hostname":   c.Hostname(),
			"pa":         os.Getenv("PROXY_API"),
			"env":        os.Getenv("ENVIRONMENT"),
		},
		"message": "",
	})

}

func convertTimestamp(timestamp string) (string, error) {
	t, err := time.Parse("Mon Jan 2 15:04:05 2006", timestamp)
	if err != nil {
		return "", err
	}
	return t.Format("02-01-2006 15:04:05"), nil
}

func getLastCommitInfo() (string, string, error) {
	cmd := exec.Command("git", "log", "-1", "--pretty=format:%h %cd", "--date=local")
	output, err := cmd.Output()
	if err != nil {
		return "", "", err
	}
	outputStr := strings.TrimSpace(string(output))
	fields := strings.Fields(outputStr)
	if len(fields) <= 2 {
		return "", "", fmt.Errorf("unexpected output format: %q", outputStr)
	}
	return fields[0], strings.Join(fields[1:], " "), nil
}

func sendEmail(to, subject, body string) error {
	from := config.LoadConfig().MailUsername
	password := config.LoadConfig().MailPassword

	// Email configuration
	smtpHost := config.LoadConfig().MailHost
	smtpPort := config.LoadConfig().MailPort
	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Send email using SMTP
	err := smtp.SendMail(smtpHost+":"+strconv.Itoa(smtpPort), auth, from, []string{to}, msg)
	if err != nil {
		return err
	}

	return nil
}

func executeCommand(commands []string, kirimEmail bool) (results []CommandResult) {
	for _, cmdStr := range commands {
		cmd := exec.Command("bash", "-c", cmdStr)

		output, err := cmd.Output()
		result := CommandResult{
			Command: cmdStr,
			Output:  string(output),
			Error:   "",
		}
		if err != nil {
			result.Error = err.Error()
		}
		results = append(results, result)
	}
	if kirimEmail {
		err := sendEmail(config.LoadConfig().NotifSystemEmail, "Command Success", "Yarn Build Finish")
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}
