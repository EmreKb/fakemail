# Fakemail

Fakemail is a Golang package designed for managing temporary email services. It provides a simple interface to interact with temp mail services like Guerrilla Mail, making it easy to generate temporary email addresses and retrieve received emails.

## Installation

```bash
go get -u github.com/EmreKb/fakemail/guerrillamail
```

## Usage

Here is an example of how to use Fakemail with Guerrilla Mail:

```go
package main

import (
	"fmt"

	"github.com/EmreKb/fakemail/guerrillamail"
)

func main() {
	mailClient := guerrillamail.New()
	emailAddress, err := mailClient.GetEmailAddress()
	if err != nil {
		panic(err) // Handle error
	}

	fmt.Printf("Email address is: %s\n", emailAddress)

	mails, err := mailClient.GetMails()
	if err != nil {
		panic(err) // Handle error
	}

	for _, mail := range mails {
		fmt.Println(mail.From)
		fmt.Println(mail.Subject)
		fmt.Println(mail.Content)
	}
}
```


## Contributing

Contributions are welcome! If you'd like to improve this package, feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Happy coding with Fakemail! 🎉
