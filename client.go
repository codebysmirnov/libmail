package libmail

import (
	"errors"
	"net/url"
	"strings"

	"gopkg.in/mail.v2"
)

// Mailer represents a structure for sending email messages.
type Mailer struct {
	sender   string  // Sender of the email.
	host     url.URL // Email server host.
	port     int     // Email server port.
	username string  // Username for authentication on the email server.
	password string  // Password for authentication on the email server.
}

// Package level variables defining common errors related to Mailer configuration.
var (
	// errEmptyMailerSender is an error indicating that the sender's email address is empty.
	errEmptyMailerSender = errors.New("empty sender")

	// errEmptyMailerHost is an error indicating that the mail server's host address is empty.
	errEmptyMailerHost = errors.New("empty host")

	// errZeroMailerPort is an error indicating that the mail server's port number is zero.
	errZeroMailerPort = errors.New("zero port")

	// errEmptyMailerUser is an error indicating that the user for Mailer authentication is empty.
	errEmptyMailerUser = errors.New("empty user")

	// errEmptyRecipients is an error indicating that the message must contain recipients.
	errEmptyRecipients = errors.New("the error occurs when recipients are not specified when sending")
)

// NewMailer creates a new instance of Mailer.
// Takes parameters for the sender, email server host, port, and password.
// Returns a pointer to Mailer.
func NewMailer(
	sender string,
	host string,
	port int,
	user string,
	password string,
) (*Mailer, error) {
	cleanedSender := strings.TrimSpace(sender)
	if cleanedSender == "" {
		return nil, errEmptyMailerSender
	}
	cleanedHost := strings.TrimSpace(host)
	if cleanedHost == "" {
		return nil, errEmptyMailerHost
	}
	parsedAPIURL, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	if port == 0 {
		return nil, errZeroMailerPort
	}
	cleanedUser := strings.TrimSpace(user)
	if cleanedUser == "" {
		return nil, errEmptyMailerUser
	}
	return &Mailer{
		sender:   sender,
		host:     *parsedAPIURL,
		port:     port,
		username: user,
		password: password,
	}, nil
}

// Send sends an email message.
// Takes a Message structure containing the recipients, subject, message text, file name, and file content.
// Returns an error if the sending process fails.
func (m *Mailer) Send(message *Message, recipients []string) error {
	if len(recipients) == 0 {
		return errEmptyRecipients
	}

	message.setSender(m.sender)
	message.setRecipients(recipients)

	// Create a new Dialer for sending the message through the email server.
	dialer := mail.NewDialer(m.host.Path, m.port, m.username, m.password)
	dialer.LocalName = m.host.Path

	// Establish the connection and send the message.
	return dialer.DialAndSend(message.msg)
}
