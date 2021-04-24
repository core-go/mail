# Mail
## Model
- Mail

## Service
- MailService

## Implementations
There are 2 implementations of MailService (You can see the samples below):
- SMTP
- SendGrid

## Installation

Please make sure to initialize a Go module before installing common-go/mail:

```shell
go get -u github.com/common-go/mail
```

Import:

```go
import "github.com/common-go/mail"
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
	"github.com/common-go/mail"
	"github.com/common-go/mail/smtp"
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
	"github.com/common-go/mail"
	"github.com/common-go/mail/sendgrid"
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
