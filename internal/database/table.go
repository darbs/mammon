package database

import (
	"context"
	//"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type Table interface {
	//MA ()
	Collection() *mongo.Collection
	AddItem(item interface{})
	GetItems(filter interface{}) []interface{}
}

type table struct {
	Name string
	Accessor ModAcc
	SliceAccessor SliceAcc
	database mongo.Database
}

func GetTable(name string) table {
	return table{
		Name: name,
		//Accessor: acc,
		//SliceAccessor: sacc,
		database: Database(),
	}
}

func (t *table) Collection() *mongo.Collection {
	return t.database.Collection(t.Name) //configurable? multiple dbs?
}

func (t *table) addItem(item interface{}) (*mongo.InsertOneResult, error) {
	return t.Collection().InsertOne(context.Background(), item)
}

type  ModAcc func() interface{}
type  SliceAcc func() []interface{}

//func New(dsn string, options ...Option) (*Flopsy, error) {
//func (t *table) getItems(filter interface{}, modelAccessor ModAcc, acc SliceAcc) ([]*interface{}, error) {
//
//	//m := make(map[string]WatchlistItem)
//	cur, err := t.Collection().Find(context.Background(), nil)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	// collection type issue
//	//main.WatchlistItem
//	//interface{}{t.model}
//	//results := getSlice(t.model)
//	//m := make([]t.model, 1)
//	//model := t.Accessor()
//	//results := make(modelAccessor(), 1)
//	defer cur.Close(context.Background())
//	for cur.Next(context.Background()) {
//		model := t.Accessor()
//		err := cur.Decode(&model)
//		if err != nil {
//			log.Fatal(err)
//			continue
//		}
//		log.Printf("%v\n", model)
//
//		//m[wli.Symbol] = wli
//		//fmt.Printf("%v\n", wli)
//	}
//	if err := cur.Err(); err != nil {
//		log.Printf("ERROR")
//	}
//
//	return []interface{}, nil
//}
