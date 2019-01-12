package internal

import (
	"time"

	"github.com/darbs/mammon/internal/database"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

const WATCH_LIST_TABLE = "Watchlist"

type WatchlistItem struct {
	Id        objectid.ObjectID `bson:"_id,omitempty" json:"Id"`
	ListName  string            `bson:"listName" json:"listName"`
	Symbol    string            `bson:"symbol" json:"symbol"`
	Name      string            `bson:"name" json:"name"`
	Country   string            `bson:"country" json:"country"`
	Date      time.Time         `bson:"date" json:"date"`
	CreatedAt time.Time         `bson:"createdAt" json:"date"`
	UpdateAt  time.Time         `bson:"updatedAt" json:"date"`
}

func (w *WatchlistItem) BeforeAdd() error {
	w.CreatedAt = time.Now()
	w.UpdateAt = time.Now()
	return nil
}

func (w *WatchlistItem) BeforeUpdate() error {
	w.UpdateAt = time.Now()
	return nil
}

func (w *WatchlistItem) GetKey() string {
	return w.Symbol
}

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
