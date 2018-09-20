package mammon

import (
	"fmt"
	"net/http"
	"astuart.co/go-robinhood"
)

const (
	apiAuth = baseUrl + "api-token-auth"
)

func dial(username string, password string) {
	creds := robinhood.Creds{Username: username, Password: password}
	token := creds.GetToken()
	fmt.Printf("TOKEN: %v\n", token)
}