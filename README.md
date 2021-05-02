# Mail
## Model
- Mail

## Service
- MailService

## Implementations
There are 2 implementations of MailService:
- SMTP
- SendGrid

## Installation
Please make sure to initialize a Go module before installing core-go/mail:

```shell
go get -u github.com/core-go/mail
```

Import:
```go
import "github.com/core-go/mail"
```

## Details
#### mail_service.go
```go
type MailService interface {
	Send(mail Mail) error
}
```

## Example of SMTP
```go
package main

import (
	"github.com/core-go/mail"
	"github.com/core-go/mail/smtp"
)

func main() {
	// Create a new smtp mail service 
	config := smtp.DialerConfig{"smtp.gmail.com", 587, "test@gmail.com", "test", true}
	mailService := smtp.NewSmtpMailSender(config)

	subject := "Your smtp demo"
	content := `Content of the email`

	mailFrom := mail.Email{Address: "peter.parker@gmail.com"}
	mailTo := []mail.Email{{Address: "mary.jane@gmail.com"}}
	mailData := mail.NewHtmlMail(mailFrom, subject, mailTo, nil, content)

	mailService.Send(*mailData)
}
```

## Example of SendGrid
```go
package main

import (
	"github.com/core-go/mail"
	"github.com/core-go/mail/sendgrid"
)

func main() {
	// Create a new sendgrid mail service 
	mailService := sendgrid.NewSendGridMailSender("xx.xxxxOQVcRKGxxxxk2KJc4g.fM7m9NIxxxxSLNOzxxxxfxF9bH4mnRrIysJA8q-xxxx")

	subject := "Your sendgrid demo"
	content := `Content of the email`

	mailFrom := mail.Email{Address: "peter.parker@gmail.com"}
	mailTo := []mail.Email{{Address: "mary.jane@gmail.com"}}
	mailData := mail.NewHtmlMail(mailFrom, subject, mailTo, nil, content)

	mailService.Send(*mailData)
}
```
