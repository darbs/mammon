package database

import (
	"fmt"
	"log"
	"time"
	"context"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type TrackWatchlist struct {
	Table Table
}

type WatchlistItem struct {
	Id       objectid.ObjectID `bson:"_id,omitempty" json:"Id"`
	ListName string            `bson:"listName" json:"ListName"`
	Symbol   string            `bson:"symbol" json:"symbol"`
	Name     string            `bson:"name" json:"name"`
	Country  string            `bson:"country" json:"country"`
	Date     time.Time         `bson:"date" json:"date"`
}

type Watchlist struct {
	table table
}

func GetWatchlist() Watchlist {
	return Watchlist{GetTable("Watchlist")}
}

func (w *Watchlist) Collection() {
	fmt.Println("Watchlist Foo")
}

// todo bulk add
func (w *Watchlist) AddItem(item interface{}) (*mongo.InsertOneResult, error) {
	fmt.Println("Watchlist Foo")
	return w.table.addItem(item)
}

func (w *WatchlistItem) GetMapKey() string {
	return w.Symbol
}

func (w *Watchlist) GetItems(filter interface{}) map[string]WatchlistItem {
	cur, err := w.table.Collection().Find(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	results := make(map[string]WatchlistItem)

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		model := WatchlistItem{}
		err := cur.Decode(&model)
		if err != nil {
			log.Fatal(err)
			continue
		}

		log.Printf("%v\n", model)
		//results = append(results, model)
		results[model.GetMapKey()] = model
	}

	if err := cur.Err(); err != nil {
		log.Printf("ERROR")
	}

	return results
}

//func (w *Watchlist) GetItems(filter interface{}) []WatchlistItem {
//	//w.table.getItems(filter, WatchlistModel, WatchlistCollection)
//	cur, err := w.table.Collection().Find(context.Background(), filter)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	results := make([]WatchlistItem, 0)
//
//	defer cur.Close(context.Background())
//	for cur.Next(context.Background()) {
//		model := WatchlistItem{}
//		err := cur.Decode(&model)
//		if err != nil {
//			log.Fatal(err)
//			continue
//		}
//		log.Printf("%v\n", model)
//		//
//		results = append(results, model)
//	}
//
//	if err := cur.Err(); err != nil {
//		log.Printf("ERROR")
//	}
//
//	return results
//}