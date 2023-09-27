package apirules

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

/*
let body = {
  To: ["test@nube-io.com"],
  Subject: "test",
  Message: "abc",
  SenderAddress: "test@nube-io.com",
  Token: "YOUR TOKEN",
};

RQL.Result = RQL.SendEmail(body);
*/

type Mail struct {
	To            []string
	Cc            []string
	Bcc           []string
	Subject       string
	Message       string
	SenderAddress string
	Token         string
}

func (inst *RQL) sendEmail(body *Mail) any {

	inst.Log(body)
	if body == nil {
		return errors.New("email body can not be empty")
	}

	to := body.To
	if len(to) <= 0 {
		return errors.New("to address can not be empty")
	}

	subject := body.Subject
	message := body.Message
	senderAddress := body.SenderAddress
	if senderAddress == "" {
		return errors.New("sender address can not be empty")
	}
	password := body.Token
	e := email.NewEmail()
	e.From = senderAddress
	e.To = to
	e.Cc = body.To
	e.Bcc = body.Bcc
	e.Subject = subject
	e.HTML = []byte(message)
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", senderAddress, password, "smtp.gmail.com"))
	if err != nil {
		return err
	}
	return fmt.Sprintf("email sent ok from: %s", senderAddress)
}

func (inst *RQL) SendEmail(body *Mail) any {
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
