package main

import (
	"fmt"
	"os"

	"github.com/darbs/mammon/internal/database"
	"github.com/darbs/mammon/internal/task"
)

func main() {
	fmt.Println("Hello, world.")

	database.Initialize(database.Connection{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Endpoint: os.Getenv("DB_ENDPOINT"),
		Port:     os.Getenv("DB_PORT"),
	})

	task.CheckWatchlist()

	/*
	client := iex.NewClient(&http.Client{})

	// Get historical data dumps available for 2016-12-12.
	histData, err := client.GetHIST(time.Date(2016, time.December, 12, 0, 0, 0, 0, time.UTC))
	if err != nil {
		panic(err)
	} else if len(histData) == 0 {
		panic(fmt.Errorf("Found %v available data feeds", len(histData)))
	}

	// Fetch the pcap dump for that date and iterate through its messages.
	resp, err := http.Get(histData[0].Link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	packetDataSource, err := iex.NewPacketDataSource(resp.Body)
	if err != nil {
		panic(err)
	}
	pcapScanner := iex.NewPcapScanner(packetDataSource)

	// Write each quote update message to stdout, in JSON format.
	enc := json.NewEncoder(os.Stdout)

	date := time.Now()
	count := 0
	for {
		msg, err := pcapScanner.NextMessage()
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		//switch  msg.(type) {
		switch msg := msg.(type) {
		case *tops.QuoteUpdateMessage:
			//t, _ := time.Parse(time.RFC3339, msg.Timestamp)
			//fmt.Println(msg)
			enc.Encode(msg)
			date = msg.Timestamp
			count++
		default:
		}
	}

	fmt.Printf("total records: %v end date: %v\n", count, date)
	*/

	//quotes, err := client.GetTOPS([]string{"AAPL", "SPY"})
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, quote := range quotes {
	//	fmt.Printf("%v: bid $%.02f (%v shares), ask $%.02f (%v shares) [as of %v]\n",
	//		quote.Symbol, quote.BidPrice, quote.BidSize,
	//		quote.AskPrice, quote.AskSize, quote.LastUpdated)
	//}
}
