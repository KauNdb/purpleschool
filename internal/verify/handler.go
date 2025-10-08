package verify

import (
	"encoding/json"
	"fmt"
	"http/configs"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type EmailIns struct {
	*configs.Config
}

const hashStr = "123"

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
		text := fmt.Sprintf("http://localhost:8081/verify/%s", hashStr)
		e.Text = []byte(text)
		e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
		e.Send("smtp.gmail.com:587", smtp.PlainAuth("", handler.Email, handler.Pass, "smtp.gmail.com"))
		json.NewEncoder(w).Encode(text)
	}
}

func (handler *EmailIns) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")
		if hash == hashStr {
			json.NewEncoder(w).Encode(true)
		} else {
			json.NewEncoder(w).Encode(false)
		}

	}
}
