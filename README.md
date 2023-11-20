# libmail

Library for sending email

The library is a wrapper over the email client, providing a more convenient and shorter use.

It contains functionality that allows you to create a message with text, subject and send it to specified mailboxes.

## How to use?

```go
package main

import (
	"log"

	"github.com/codebysmirnov/libmail"
)

func main() {
	mailer, err := libmail.NewMailer("mailsender@example.com", "smtp.server.io", 8080, "smtp.username", "smtp.userpassword")
	if err != nil {
		log.Fatalln(err)
	}

	msg := libmail.NewMessage("some subject", "some text")
	recipients := []string{"first@exmple.com", "second@example.com"}
	err = mailer.Send(msg, recipients)
	if err != nil {
		log.Fatalln(err)
	}
}
```
