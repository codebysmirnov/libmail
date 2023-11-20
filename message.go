package libmail

import (
	"bytes"
	"errors"
	"strings"

	"gopkg.in/mail.v2"
)

// Message represents a structure for an email message.
type Message struct {
	msg *mail.Message // Internal representation of the email message.
}

// NewMessage creates a new Message instance for sending an email.
// Takes sender, subject, text of the message, and a list of recipients.
// Returns a Message and an error if sender is empty or there are no recipients.
func NewMessage(sub, text string) *Message {
	var m Message
	m.msg = mail.NewMessage(mail.SetEncoding(mail.Base64))
	m.msg.SetHeaders(map[string][]string{"Subject": {sub}})
	m.msg.SetBody("text/plain", text, mail.SetPartEncoding(mail.Unencoded))
	return &m
}

// File represents a structure for a file attachment in an email.
type File struct {
	Name   string        // Name of the file.
	Reader *bytes.Reader // Reader for reading the file content.
}

var (
	errEmptyFileName = errors.New("empty file name")
	errEmptyFileBody = errors.New("file is empty")
)

// NewFile creates a new File instance.
// Takes a name for the file and the content as a byte slice.
// Returns a File and an error if the name is empty or the content is empty.
func NewFile(name string, buf []byte) (File, error) {
	cleanedName := strings.TrimSpace(name)
	if cleanedName == "" {
		return File{}, errEmptyFileName
	}
	if len(buf) == 0 {
		return File{}, errEmptyFileBody
	}
	return File{
		Name:   name,
		Reader: bytes.NewReader(buf),
	}, nil
}

// IncludeFile attaches a file to the email message.
func (m *Message) IncludeFile(f File) {
	m.msg.AttachReader(f.Name, f.Reader)
}

// setSender sets the sender's address for the email message.
func (m *Message) setSender(sender string) {
	m.msg.SetHeaders(map[string][]string{
		"From": {sender},
	})
}

// setRecipients sets the recipients' addresses for the email message.
func (m *Message) setRecipients(recipients []string) {
	m.msg.SetHeaders(map[string][]string{
		"To": recipients,
	})
}
