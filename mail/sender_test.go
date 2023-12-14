package mail

import (
	"api-go-fit-tracker/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	config, err := config.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	subject := "Just Test"
	content := `
	<h1>Hello Guy</h1>
	<p>Just Test </p>
	`
	to := []string{"sene300@gmail.com"}

	err = sender.SendEmail(subject, content, to, nil, nil, nil)
}
