package mail

type SimpleMailService interface {
	Send(mail SimpleMail) error
}
