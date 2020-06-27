package mail

type MailService interface {
	Send(mail Mail) error
}
