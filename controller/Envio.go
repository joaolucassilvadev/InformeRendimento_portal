package controller

/*
acredito que já é feito isso
import (
	"bytes"
	"io/ioutil"
	"log"
	"mime"
	"net/smtp"
	"path/filepath"
)

func EnvioEmail() {
	f := "jls.silva@discente.ufma.br"
	from := "js0454261@gmail.com"
	to := f
	subject := "Teste de e-mail com anexo"
	body := "Corpo do e-mail"
	d := "05895613322"
	// Caminho do arquivo PDF a ser anexado
	pdfPath := "C:/Users/joao.lucas/email/paginas/" + d + ".pdf"

	// Configuração do SMTP
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpUsername := "js0454261@gmail.com"
	smtpPassword := "lhts rsjs onnt wyme"

	// Autenticação SMTP
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Cria um corpo multipart
	msg := createMultipartMessage(from, to, subject, body)

	// Adiciona o anexo ao corpo multipart
	err := attachFile(msg, pdfPath)
	if err != nil {
		log.Fatalf("Erro ao anexar arquivo: %s", err)
	}

	// Converte o corpo multipart para bytes
	emailBytes := msg.Bytes()

	// Envia o e-mail
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, emailBytes)
	if err != nil {
		log.Fatalf("Erro ao enviar e-mail: %s", err)
	}

	log.Println("E-mail enviado com sucesso")
}

// Cria um corpo multipart para o e-mail
func createMultipartMessage(from, to, subject, body string) *bytes.Buffer {
	var msg bytes.Buffer

	headers := map[string]string{
		"From":    from,
		"To":      to,
		"Subject": subject,
	}

	for key, value := range headers {
		msg.WriteString(key + ": " + value + "\r\n")
	}
	msg.WriteString("MIME-version: 1.0;\nContent-Type: multipart/mixed; boundary=\"boundary\"\r\n")
	msg.WriteString("\r\n")

	msg.WriteString("--boundary\r\n")
	msg.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	msg.WriteString("\r\n")
	msg.WriteString(body + "\r\n")

	return &msg
}

// Anexa um arquivo ao corpo multipart do e-mail
func attachFile(msg *bytes.Buffer, filePath string) error {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	fileName := filepath.Base(filePath)
	mimeType := mime.TypeByExtension(filepath.Ext(filePath))

	msg.WriteString("\r\n--boundary\r\n")
	msg.WriteString("Content-Disposition: attachment; filename=\"" + fileName + "\"\r\n")
	msg.WriteString("Content-Type: " + mimeType + "\r\n")
	msg.WriteString("\r\n")

	msg.Write(fileContent)

	return nil
}
*/
