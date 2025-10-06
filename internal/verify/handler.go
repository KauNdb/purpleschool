package verify

import (
	"fmt"
	"http/configs"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type EmailIns struct {
	*configs.Config
}

func NewEmail(router *http.ServeMux, config *configs.Config) {
	emailIns := &EmailIns{
		Config: config,
	}
	router.HandleFunc("POST /send", emailIns.Send())
	router.HandleFunc("GET /verify/{hash}", emailIns.Verify())
}

func (handler *EmailIns) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		e := email.NewEmail()
		FromMsg := fmt.Sprintf("Jordan Wright <%s>", handler.Email)
		e.From = FromMsg
		e.To = []string{handler.Addr}
		e.Subject = "Awesome Subject"
		e.Text = []byte("Text Body is, of course, supported!")
		e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
		e.Send("smtp.gmail.com:587", smtp.PlainAuth("", handler.Email, handler.Pass, "smtp.gmail.com"))
	}
}

func (handler *EmailIns) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
	}
}
