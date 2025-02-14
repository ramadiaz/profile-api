package dto

type EmailRequest struct {
	Email   string
	Subject string
	Body    string
}

type EmailIncognites struct {
	Subject       string
	RecipientName string
	MessageBody   string
	SecurityLevel string
	SentDate      string
	MessageID     string
}