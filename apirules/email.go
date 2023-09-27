package apirules

import (
	"encoding/json"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type Mail struct {
	To            []string
	Cc            []string
	Bcc           []string
	Subject       string
	Message       string
	SenderAddress string
	Token         string
}

func (inst *RQL) sendEmail(body *Mail) error {
	parsed, err := emailBody(body)
	if err != nil {
		return err
	}
	to := parsed.To
	subject := parsed.Subject
	message := parsed.Message
	senderAddress := parsed.SenderAddress
	password := parsed.Token
	e := email.NewEmail()
	e.From = senderAddress
	e.To = to
	e.Cc = parsed.To
	e.Bcc = parsed.Bcc
	e.Subject = subject
	e.HTML = []byte(message)
	return e.Send("smtp.gmail.com:587", smtp.PlainAuth("", senderAddress, password, "smtp.gmail.com"))
}

func (inst *RQL) SendEmail(body *Mail) error {
	return inst.sendEmail(body)

}

func emailBody(body any) (*Mail, error) {
	result := &Mail{}
	dbByte, err := json.Marshal(body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(dbByte, &result)
	return result, err
}
