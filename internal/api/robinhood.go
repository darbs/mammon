package api

import (
	"fmt"
	"time"

	//"time"

	"github.com/darbs/mammon/internal/database"
	log "github.com/sirupsen/logrus"

	"astuart.co/go-robinhood"
)

var logger *log.Entry

type rhConnection struct {
	connection *robinhood.Client
}

/*
Option?:
Refactor this into a generic Dial(params interface{}) method
that reflects the type of param and return the api specific type
 */
func RobinhoodDial(username string, password string) (*rhConnection, error) {
	cli, err := robinhood.Dial(&robinhood.OAuth{
		Username: username,
		Password: password,
	})

	if err != nil {
		fmt.Printf("ERR: %v\n", err)
		return nil, ErrApiConnect
	}

	return &rhConnection{
		connection: cli,
	}, nil
}

// TODO finish building out rh specific connection as well as genericising the api interface
func (rhc *rhConnection) GetWatchlist() []database.WatchlistItem {
	items := make([]database.WatchlistItem, 0)
	watchlists, err := rhc.connection.GetWatchlists()
	if err != nil {
		fmt.Printf("ERR RETRIEVING WATCHLIST: %v\n", err)
	}

	for index, watchlist := range watchlists {
		fmt.Printf("%v %v\n", index, watchlist)
		tickers, err := watchlist.GetInstruments()
		if err != nil {
			fmt.Errorf("Error retrieving tickers for watchlist %v: %v\n", watchlist.Name, err)
			continue
		}

		for _, ticker := range tickers {
			//fmt.Printf("watchlist item %v\n", ticker)
			// todo set/collection?
			items = append(items, database.WatchlistItem{
				Symbol:  ticker.Symbol,
				Name:    ticker.Name,
				Country: ticker.Country,
				Date:    time.Now().UTC(),
			})
		}
	}

	return items
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	logger = log.WithFields(log.Fields{
		"subject": "table",
		"name":    "Robinhood",
	})

	logger.Info("foobar")
}