package emails

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"profile-api/emails/dto"
	"profile-api/emails/templates"
	"profile-api/pkg/exceptions"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

func SendEmail(data dto.EmailRequest) *exceptions.Exception {
	email := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	server := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")

	i, err := strconv.Atoi(smtpPort)
	if err != nil {
		return exceptions.NewException(http.StatusInternalServerError, exceptions.ErrInternalServer)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", data.Email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", data.Body)

	d := gomail.NewDialer(server, i, email, password)

	if err := d.DialAndSend(m); err != nil {
		return exceptions.NewException(http.StatusBadGateway, exceptions.ErrEmailSendFailed)
	}

	fmt.Println("Email sent successfully to: ", data.Email)

	return nil
}

func SendIncognitoEmail(data dto.EmailIncognites) *exceptions.Exception {
	tmpl, err := template.New("email").Parse(templates.IncognitoEmailTemplate)
	if err != nil {
		return exceptions.NewException(http.StatusInternalServerError, exceptions.ErrFailedToParseTemplate)
	}

	if data.SecurityLevel == "" {
		data.SecurityLevel = "Confidential"
	}
	if data.SentDate == "" {
		data.SentDate = time.Now().Format("January 02, 2006 15:04 MST")
	}
	if data.MessageID == "" {
		data.MessageID = "MSG-" + time.Now().Format("20060102150405")
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, data); err != nil {
		return exceptions.NewException(http.StatusInternalServerError, exceptions.ErrFailedToExecuteTemplate)
	}

	err = SendEmail(dto.EmailRequest{
		Email:   "ramadiaz221@gmail.com",
		Subject: "Incognito Message",
		Body:    buffer.String(),
	})

	return nil
}
