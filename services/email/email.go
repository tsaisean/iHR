package email

import (
	"fmt"
	"iHR/config"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailService struct {
	appURL string
	senderName   string
	config config.Email
}

func NewEmailService(appURL string, senderName string, config config.Email) *EmailService {
	return &EmailService{
		appURL:     appURL,
		senderName: senderName,
		config:     config,
	}
}

func (s *EmailService) SendPasswordResetEmail(receiverName string, receiverEmail string, token string) error {

	resetLink := fmt.Sprintf("%s/reset-password?token=%s", s.appURL, token)
	htmlContent := fmt.Sprintf(`
		<html>
			<body>
				<h2>Password Reset Request</h2>
				<p>You have requested to reset your password. Click the link below to proceed:</p>
				<p><a href="%s">Reset Password</a></p>
				<p>If you didn't request this, please ignore this email.</p>
				<p>This link will expire in 24 hours.</p>
			</body>
		</html>
	`, resetLink)

	from := mail.NewEmail(s.senderName, s.config.SenderEmail)
	subject := "Password Reset Request"
	to := mail.NewEmail(receiverName, receiverEmail)
	plainTextContent := "and easy to do anywhere, even with Go"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(s.config.SendgridKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}

	return nil
}
