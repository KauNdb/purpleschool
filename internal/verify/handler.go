package verify

import (
	"encoding/json"
	"fmt"
	"http/configs"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
	"golang.org/x/crypto/bcrypt"
)

type EmailIns struct {
	*configs.Config
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JsonUser struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
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
		var user User
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		defer req.Body.Close()

		e := email.NewEmail()
		FromMsg := fmt.Sprintf("Jordan Wright <%s>", user.Email)
		e.From = FromMsg
		e.To = []string{handler.Addr}
		e.Subject = "Awesome Subject"
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		jsonUser := JsonUser{
			Email: user.Email,
			Hash:  string(hash),
		}

		// сериализуем в JSON
		file, err := os.Create("user.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(jsonUser); err != nil {
			log.Fatal(err)
		}

		text := fmt.Sprintf("http://localhost:8081/verify/%s", string(hash))
		e.Text = []byte(text)
		e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
		e.Send("smtp.gmail.com:587", smtp.PlainAuth("", user.Email, user.Password, "smtp.gmail.com"))
		json.NewEncoder(w).Encode(text)
	}
}

func (handler *EmailIns) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")
		data, err := os.ReadFile("user.json")
		if err != nil {
			log.Fatal(err)
		}

		var loaded JsonUser
		if err := json.Unmarshal(data, &loaded); err != nil {
			log.Fatal(err)
		}
		if loaded.Hash == hash {
			json.NewEncoder(w).Encode(true)
		} else {
			os.Remove("user.json")
			json.NewEncoder(w).Encode(false)
		}
	}
}
