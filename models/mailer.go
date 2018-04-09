package models

import (
	"bytes"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"
)

type Mailer struct {
	Sender     mail.Address
	Receiver   mail.Address
	Server     string
	PortNumber int
	Auth       smtp.Auth
	Data       *SnippetInfo
}

var SmtpMailer Mailer

const tpl = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>{{ .Title }}</title>
  <style>
    code {
    display: block;
    white-space: pre-wrap;
    background: hsl(220, 80%, 90%);
  }
</style>
</head>
  <p><b>{{ .Owner }}</b> from Code Directour shared the following snippet with you.<p>
  <br/>
  <p><b>Language:</b> {{ .Language }}</p>
  <code>
    {{ .Code }}
  </code>
  <p><b>References:</b> {{ .References }}</p>
<body>
</body>
</html>`

func NewMailer(server string, port int, sender, receiver mail.Address, userName, secret string, s *SnippetInfo) {
	SmtpMailer = Mailer{
		Sender:     sender,
		Receiver:   receiver,
		Server:     server,
		PortNumber: port,
		Auth:       smtp.PlainAuth("", userName, secret, server),
		Data:       s,
	}
}

var header map[string]string

func toString(m map[string]string) string {
	var concat string
	for k, v := range m {
		concat += fmt.Sprintf("%s:%s\r\n", k, v)
	}
	return concat
}

func (m *Mailer) MakeHeader() string {
	header = make(map[string]string)
	header["MIME-Version"] = "1.0"
	header["From"] = m.Sender.String()
	header["To"] = m.Receiver.String()
	header["Subject"] = fmt.Sprintf("%s Snippet Shared: %s", m.Data.Language, m.Data.Title)
	header["Content-type"] = "text/html"
	return toString(header)
}

func (m *Mailer) MailBody() []byte {

	t, err := template.New("webpage").Parse(tpl)
	var buff bytes.Buffer
	err = t.Execute(&buff, m.Data)
	if err != nil {
		return make([]byte, 0)
	}
	return []byte(m.MakeHeader() + buff.String())
}

func (m *Mailer) ServerName() string {
	return fmt.Sprintf("%s:%d", m.Server, m.PortNumber)
}

func (m *Mailer) SendMail() error {
	return smtp.SendMail(m.ServerName(), m.Auth, m.Sender.Address, []string{m.Receiver.Address}, m.MailBody())
}
