package Utils

import (
	b64 "encoding/base64"
	"fmt"
	"log"
	"net/smtp"
	"math/rand"
	"time"
	"golang.org/x/crypto/bcrypt"
)
const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func EncodeBase64(input string) string {
	sEnc := b64.StdEncoding.EncodeToString([]byte(input))
	fmt.Println(sEnc)
	return sEnc
}

func DecodeBase64(input string) string {

    sDec, _ := b64.StdEncoding.DecodeString(input)
    fmt.Println(string(sDec))
    fmt.Println()
	return string(sDec)
}

func SendEmail(body string, receiver string) {
	from := "distribuidas2020@gmail.com"
	pass := "Admin2020"
	to := receiver

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Nueva contraseña de ingreso\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	
	log.Print("email sent")
}


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

// func ConvertTimestampToDate(input string) string {

// 	i, err := strconv.ParseInt(input, 10, 64)
//     if err != nil {
// 		panic(err)
// 	}
// 	tm := time.Unix(i, 0)
// 	fmt.Println(tm)
// 	return string(tm.Format("2006-01-02 15:04:05"))

// }
