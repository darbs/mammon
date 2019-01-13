package internal

import (
	"fmt"
	"time"

	"astuart.co/go-robinhood"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

type rhConnection struct {
	connection *robinhood.Client
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	logger = log.WithFields(log.Fields{
		"subject": "externalApi",
		"name":    "robinhood",
	})
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
		return nil, err
	}

	return &rhConnection{
		connection: cli,
	}, nil
}

// TODO finish building out rh specific connection as well as genericising the api interface
func (rhc *rhConnection) GetWatchlist() ([]WatchlistItem, error) {
	items := make([]WatchlistItem, 0)
	watchlists, err := rhc.connection.GetWatchlists()

	if err != nil {
		logger.Error(err)
		return items, err
	}

	for index, watchlist := range watchlists {
		fmt.Printf("%v %v\n", index, watchlist)
		tickers, err := watchlist.GetInstruments()
		if err != nil {
			logger.Warnf("Error retrieving tickers for watchlist %v: %v\n", watchlist.Name, err)
			continue
		}

		for _, ticker := range tickers {
			// todo set/collection?
			items = append(items, WatchlistItem{
				Symbol:  ticker.Symbol,
				Name:    ticker.Name,
				Country: ticker.Country,
				Date:    time.Now().UTC(),
			})
		}
	}

	return items, err
}