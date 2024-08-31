package email

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

func isValidEmail(email string) bool {
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)

	return re.MatchString(email)
}

const myMail = "kdesp2003@gmail.com"

func SendEmail(name, email, body string, resultChan chan error) {
	if strings.TrimSpace(name) == "" {
		resultChan <- fmt.Errorf("Name is empty")
	}
	if strings.TrimSpace(email) == "" {
		resultChan <- fmt.Errorf("Email is empty")
	}
	if !isValidEmail(email) {
		resultChan <- fmt.Errorf("Email is invalid")
	}
	if strings.TrimSpace(body) == "" {
		resultChan <- fmt.Errorf("Body is empty")
	}

	err := godotenv.Load()
	if err != nil {
		resultChan <- err
	}
	pass := os.Getenv("GOOGLE_APP_PASSWORD")

	auth := smtp.PlainAuth(
		"", 
		myMail, 
		pass, 
		"smtp.gmail.com",
	)

	to := []string{myMail}

	msg := []byte(
		"To: " + myMail + "\r\n" +
		"Subject: Ssh portfolio message by " + name + "(" + email + ")\r\n" +
		"\r\n" +
		body + "\r\n",
	)

	err = smtp.SendMail("smtp.gmail.com:587", auth, myMail, to, msg)

	if err != nil {
		log.Fatal(err)
		resultChan <- err
	}

	resultChan <- nil
}
