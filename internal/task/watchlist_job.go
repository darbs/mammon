package task

import (
	"os"

	"github.com/darbs/mammon/internal"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

func init() {
	log.SetFormatter(&log.TextFormatter{})
	logger = log.WithFields(log.Fields{
		"subject": "job",
	})
}

func CheckWatchlist() error {
	rhapi, err := internal.RobinhoodDial(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		panic(err)
	}

	apiItems, err := rhapi.GetWatchlist()
	logger.Debugf("Received %v watchlist items", len(apiItems))
	internal.SetWatchlist(apiItems)

	return err
}
