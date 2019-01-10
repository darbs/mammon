package main

import (
	"fmt"
	"os"
	//"time"

	"github.com/darbs/mammon/internal/api"
	"github.com/darbs/mammon/internal/database"
)

func main() {
	fmt.Println("Hello, world.")

	database.Initialize(database.Connection{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Endpoint: os.Getenv("DB_ENDPOINT"),
		Port: os.Getenv("DB_PORT"),
	})

	db := database.Database()

	//w := database.WatchlistHistory()
	//
	//w.Table.Foo()

	//w.Foo()

	fmt.Printf("%v\n", db)
	rhapi, err  := api.RobinhoodDial(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		panic(err)
	}

	apiItems := rhapi.GetWatchlist()

	//dbItems := make([]database.WatchlistItem, 0)
	//dbItems := make(map[string]database.WatchlistItem, 0)
	dbItems := make(map[string]database.WatchlistItem, 0)
	watchlistTable := database.Table(database.WATCH_LIST_TABLE)
	err = watchlistTable.GetItems(nil, &dbItems)
	if err != nil {
		panic(err)
	}

	fmt.Println(apiItems)
	fmt.Println(dbItems)

	apiItemMap := make(map[string]database.WatchlistItem, 0)
	for _, wli := range apiItems {
		fmt.Printf("%v %v %v \n",wli.Country, wli.Symbol, wli.Name)
		apiItemMap[wli.GetKey()] = wli

		//if _, ok := exisingItems[ticker.Symbol]; ok {
		//	fmt.Printf("Existing: %v\n", ticker.Symbol)
		//	delete(exisingItems, ticker.Symbol)
		//} else {
		//	model := database.WatchlistItem{
		//		Symbol: ticker.Symbol,
		//		Name: ticker.Name,
		//		Country: ticker.Country,
		//		Date: time.Now().UTC(),
		//	}

			//exisingItems[model.GetMapKey()] = model
		//}

		//_, err := watchlistTable.AddItem(&wli)
		//if err != nil {
		//	//logger.Error(err)
		//	panic("AHHHHHHH")
		//}
	}

	//dbItemMap := make(map[string]database.WatchlistItem, 0)
	//for _, wli := range dbItems {
	//	fmt.Printf("%v %v %v \n", wli.Country, wli.Symbol, wli.Name)
	//	dbItemMap[wli.GetKey()] = wli
	//}

	//for _, wli := range apiItems {
	//	if _, ok := dbItemMap[wli.GetKey()]; !ok {
	//		fmt.Printf("Existing: %v\n", ticker.Symbol)
	//		delete(exisingItems, ticker.Symbol)
	//	}
	//}

	//fmt.Printf("NEED TO ADD: %v\n", exisingItems)
	//for _, value := range exisingItems {
	//	res, err := w.AddItem(value)
	//
	//	if err != nil {
	//		fmt.Errorf("Error add watchlist item %v: %v\n", watchlist.Name, err)
	//		continue
	//	}
	//
	//	fmt.Printf("Result: %v\n", res)
	//}

	//////////////////////

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
