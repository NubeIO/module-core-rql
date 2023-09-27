package apirules

import (
	"fmt"
	"testing"
)

func TestRQL_sendEmail(t *testing.T) {
	r := RQL{}
	e := &Mail{
		To:            []string{"aidan@TEST.com.au"},
		Cc:            nil,
		Bcc:           nil,
		Subject:       "test",
		Message:       "test",
		SenderAddress: "ap@TEST.com",
		Token:         "", // see to make token // https://devanswe.rs/create-application-specific-password-gmail/
	}
	err := r.sendEmail(e)
	fmt.Println(err)
	if err != nil {
		return
	}
}
