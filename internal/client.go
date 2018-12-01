package internal

import (
	"fmt"
	"astuart.co/go-robinhood"
)

const (
	//apiAuth = baseUrl + "api-token-auth"
)

func Dial(username string, password string) {
	//fmt.Printf("TOKEN: %v CREDS: %v\n", token, creds.Values().Encode())

	cli, err := robinhood.Dial(&robinhood.OAuth{
		Username: username,
		Password: password,
	})

	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}

	watchlists, err := cli.GetWatchlists()
	if err != nil {
		fmt.Printf("ERR RETRIEVING WATCHLIST: %v\n", err)
	}

	//fmt.Printf("%v\n", instruments)

	for index, watchlist := range watchlists {
		fmt.Printf("%v %v\n", index, watchlist)
		tickers, err := watchlist.GetInstruments()
		if err != nil {
			fmt.Errorf("Error retrieving tickers for watchlist %v: %v\n", watchlist.Name, err)
			continue
		}

		fmt.Printf("WATCHLIST %v\n", watchlist.Name)
		for _, ticker := range tickers {
			fmt.Printf("%v %v %v %v %v\n", ticker.Country, ticker.Symbol, ticker.Name, ticker.ID, ticker.BloombergUnique)
		}

	}
}