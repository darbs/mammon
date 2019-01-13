package internal

import (
	"time"

	"github.com/darbs/mammon/internal/database"
	"github.com/mongodb/mongo-go-driver/bson"
)

const WATCH_LIST_TABLE = "Watchlist"

type WatchListTable struct {
	*database.Table
}

func (t *WatchListTable) RemoveItemsBySymbol(symbols *[]string) (int64, error) {
	query := bson.D{
		{"symbol",
			bson.D{
				{"$in", *symbols},
			},
		},
	}

	//logger.Warnf("query: %v\n", query)
	return t.RemoveItems(query)
}

//func (t *WatchListTable) AddItems(items []WatchlistItem) (error) {
//	docs := make([]interface{}, len(items))
//	for i := 0; i < len(items); i++ {
//		d := items[i]
//		docs[i] = &d
//	}
//	_, err := t.Collection().InsertMany(context.Background(), docs)
//	return err
//}

func GetWatchlistTable() WatchListTable {
	table := database.GetTable(WATCH_LIST_TABLE)
	return WatchListTable{&table}
}

/*

 */
func SetWatchlist(items []WatchlistItem) error {
	table := GetWatchlistTable()
	creationTime := time.Now()
	existing := make(map[string]WatchlistItem, 0)

	dbItems := make(map[string]WatchlistItem, 0)
	err := table.GetItems(nil, &dbItems)
	if err != nil {
		return err
	}

	// check which items need to get added
	itemsToAdd := make([]WatchlistItem, 0)
	for _, wli := range items {
		if _, ok := dbItems[wli.Symbol]; !ok {
			wli.UpdateAt = creationTime
			wli.CreatedAt = creationTime
			itemsToAdd = append(itemsToAdd, wli)
		} else {
			existing[wli.Symbol] = wli
		}
	}

	// check which items need to get deleted
	itemsToDelete := make([]string, 0)
	for _, wli := range dbItems {
		if _, ok := existing[wli.Symbol]; !ok {
			itemsToDelete = append(itemsToDelete, wli.Symbol)
		}
	}

	logger.Printf("itemsToAdd %v", itemsToAdd)
	addResult, err := table.AddItems(&itemsToAdd)
	logger.Printf("Add result: %v", addResult)
	if err != nil {
		return err
	}

	logger.Printf("itemsToDelete %v", itemsToDelete)
	remResult, err := table.RemoveItemsBySymbol(&itemsToDelete)
	logger.Printf("Remove result: %v", remResult)
	if err != nil {
		return err
	}

	return nil
}
