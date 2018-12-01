package database

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var instance *db
var once sync.Once

type db struct {
	initialized bool
	client     mongo.Client
}

type Connection struct {
	Username string
	Password string
	Endpoint string
	Port     string
}

func (c *Connection) getConnectionString() string {
	//"mongodb://foo:bar@localhost:27017"
	return "mongodb://" + c.Username + ":" + c.Password + "@" + c.Endpoint + ":" + c.Port
}

func getConnection(params Connection) mongo.Client {
	url := params.getConnectionString()
	client, err := mongo.NewClient(url)
	if err != nil {
		log.Fatal(err)
		panic("FAILED TO INITIALIZE CLIENT CONNECTION TO DB")
		return *client
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		panic("CLIENT FAILED TO CONNECT TO DB")
		return *client
	}

	return *client
}



func Initialize(params Connection) error {
	once.Do(func() {
		//if instance.initialized { return }

		client := getConnection(params)
		instance = &db{client: client}
	})

	return nil
}

func Database() mongo.Database {
	return *instance.client.Database("Mammon") //configurable? multiple dbs?
}
