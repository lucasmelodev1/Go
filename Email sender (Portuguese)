package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	emailList := []string{}

	autor := mail.NewEmail("Lucas Gabriel", "lucasgabrielmelorodrigues@gmail.com")
	titulo := "Essa API é muito legal!"
	textocru := `Consegui fazer esse email de forma bem simples!
	E a documentação deles é bem clara!
	O site também é muito bonito.`
	textohtml := `Consegui fazer esse email de forma bem simples!
	<strong>E a documentação deles é bem clara</strong>!
	O site também é muito bonito.`
	cliente := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	os.Setenv("SENDGRID_API_KEY", "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")
	for _, email := range emailList {
		remetente := mail.NewEmail("Cliente", email)
		envio := mail.NewSingleEmail(autor, titulo, remetente, textocru, textohtml)
		response, err := cliente.Send(envio)

		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
		}
	}

}
