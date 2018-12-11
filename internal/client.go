package internal

import (
	"fmt"
	"time"

	//"time"

	"astuart.co/go-robinhood"
	"github.com/darbs/mammon/internal/database"
	//"github.com/mongodb/mongo-go-driver/mongo"

	//"github.com/mongodb/mongo-go-driver/bson"
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

	w := database.GetWatchlist()
	exisingItems := w.GetItems(nil)
	fmt.Printf("%v\n", exisingItems)

	for index, watchlist := range watchlists {
		fmt.Printf("%v %v\n", index, watchlist)
		tickers, err := watchlist.GetInstruments()
		if err != nil {
			fmt.Errorf("Error retrieving tickers for watchlist %v: %v\n", watchlist.Name, err)
			continue
		}

		fmt.Printf("WATCHLIST %v\n", watchlist.Name)
		for _, ticker := range tickers {
			fmt.Printf("%v %v %v %v %v %v\n", ticker.Market, ticker.Country, ticker.Symbol, ticker.Name, ticker.ID, ticker.BloombergUnique)

			if _, ok := exisingItems[ticker.Symbol]; ok {
				delete(exisingItems, ticker.Symbol)
			} else {
				model := database.WatchlistItem{
					Symbol: ticker.Symbol,
					Name: ticker.Name,
					Country: ticker.Country,
					Date: time.Now().UTC(),
				}
				exisingItems[model.GetMapKey()] = model
			}
		}

		fmt.Printf("NEED TO ADD: %v\n", exisingItems)
		for _, value := range exisingItems {
			res, err := w.AddItem(value)

			if err != nil {
				fmt.Errorf("Error add watchlist item %v: %v\n", watchlist.Name, err)
				continue
			}

			fmt.Printf("Result: %v\n", res)
		}
	}
}