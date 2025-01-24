package guerrillamail

import (
	"github.com/EmreKb/fakemail/http_client"
)

const (
	HOST = "api.guerrillamail.com"
	PATH = "ajax.php"

	FUNCTION_GET_EMAIL_ADDRESS = "get_email_address"
	FUNCTION_CHECK_EMAIL       = "check_email"
)

type MailClient struct {
	client *http_client.HttpClient
}

func New(opts ...http_client.ClientOpts) *MailClient {
	client := http_client.New(opts...)

	return &MailClient{
		client: client,
	}
}

type getEmailAddress struct {
	EmailAddr string `json:"email_addr"`
}

func (c *MailClient) GetEmailAddress() (string, error) {
	url := http_client.GetUrl(
		HOST,
		http_client.WithPath(PATH),
		http_client.WithQueries(map[string]string{"f": FUNCTION_GET_EMAIL_ADDRESS}),
	)
	var data getEmailAddress
	if err := c.client.Get(url, &data); err != nil {
		return "", err
	}

	return data.EmailAddr, nil
}

type Mail struct {
	From    string
	Subject string
	Content string
}

type email struct {
	MailFrom    string `json:"mail_from"`
	MailSubject string `json:"mail_subject"`
	MailContent string `json:"mail_excerpt"`
}

type checkEmail struct {
	List []email `json:"list"`
}

func (c *MailClient) GetMails() ([]Mail, error) {
	url := http_client.GetUrl(
		HOST,
		http_client.WithPath(PATH),
		http_client.WithQueries(map[string]string{"f": FUNCTION_CHECK_EMAIL, "seq": "0"}),
	)
	var data checkEmail
	if err := c.client.Get(url, &data); err != nil {
		return nil, err
	}

	mails := make([]Mail, len(data.List))
	for i, m := range data.List {
		mails[i] = Mail{
			From:    m.MailFrom,
			Subject: m.MailSubject,
			Content: m.MailContent,
		}
	}

	return mails, nil
}
