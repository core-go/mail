# Mail
## Model
- Mail

## Service
- MailService

## Implementations
There are 2 implementations of MailService (You can see the samples below):
- SMTP at [common-go/smtp](https://github.com/common-go/smtp)
- SendGrid at [common-go/sendgrid](https://github.com/common-go/sendgrid)

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

## Example of [SMTP](https://github.com/common-go/smtp)
```go
package main

import (
	"github.com/common-go/mail"
	"github.com/common-go/smtp"
)

func main() {
	// Create a new smtp mail service 
	config := smtp.DialerConfig{"smtp.gmail.com", 587, "test@gmail.com", "test", true}
	mailService := smtp.NewSmtpMailService(config)

	subject := "Your smtp demo"
	content := `Content of the email`

	mailFrom := mail.Email{Address: "peter.parker@gmail.com"}
	mailTo := []mail.Email{{Address: "mary.jane@gmail.com"}}
	mailData := mail.NewHtmlMail(mailFrom, subject, mailTo, nil, content)

	mailService.Send(*mailData)
}
```

## Example of [SendGrid](https://github.com/common-go/smtp)
```go
package main

import (
	"github.com/common-go/mail"
	"github.com/common-go/sendgrid"
)

func main() {
	// Create a new sendgrid mail service 
	mailService := sendgrid.NewSendGridMailService("xx.xxxxOQVcRKGxxxxk2KJc4g.fM7m9NIxxxxSLNOzxxxxfxF9bH4mnRrIysJA8q-xxxx")

	subject := "Your sendgrid demo"
	content := `Content of the email`

	mailFrom := mail.Email{Address: "peter.parker@gmail.com"}
	mailTo := []mail.Email{{Address: "mary.jane@gmail.com"}}
	mailData := mail.NewHtmlMail(mailFrom, subject, mailTo, nil, content)

	mailService.Send(*mailData)
}
```
