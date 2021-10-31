# Mail
![Email](https://camo.githubusercontent.com/1a239ae784d3f8c33517b2b4b0860f6e438432c6589191072ae773204918fc39/68747470733a2f2f63646e2d696d616765732d312e6d656469756d2e636f6d2f6d61782f3830302f312a374a46344f47397a427973754a71724f4958324a76672e706e67)
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
