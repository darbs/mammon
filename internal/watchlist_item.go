package internal

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

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
