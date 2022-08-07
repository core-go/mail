# Mail
![Email](https://camo.githubusercontent.com/0fcd0826eea9b9883077ac1674e45c4eafa17e7abb02f8d2659fb30bc2b084e5/68747470733a2f2f63646e2d696d616765732d312e6d656469756d2e636f6d2f6d61782f3830302f312a2d6c486a7872355a4d6b4b634c6961746776364731672e706e67)
- The sample is [go-authentication](https://github.com/project-samples/go-authentication)

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
