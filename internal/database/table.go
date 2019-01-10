package database

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

type Model interface {
	GetKey() string
	PreAdd() error
	PreUpdate() error
}

type CollectionObject interface {
	GetKey() string
}

type DbCallbackBeforeAdd interface {
	BeforeAdd() error
}

type DbCallbackBeforeUpdate interface {
	BeforeUpdate() error
}

/*
Todo
have this as part of the attribute structure for desired models akin to gorm
 */
type table struct {
	Name     string
	database mongo.Database
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	logger = log.WithFields(log.Fields{
		"subject": "table",
		"name":   "main",
	})
}

func Table(name string) table {
	return table{
		Name:     name,
		database: Database(),
	}
}

func (t *table) Collection() *mongo.Collection {
	return t.database.Collection(t.Name) //configurable? multiple dbs?
}

func (t *table) AddItem(item interface{}) (*mongo.InsertOneResult, error) {
	if reflect.ValueOf(item).Kind() != reflect.Ptr {
		return nil, ErrExpectingPointer
	}

	if obj, ok := item.(DbCallbackBeforeAdd); ok {
		obj.BeforeAdd()
	}

	return t.Collection().InsertOne(context.Background(), item)
}

func (t *table) GetItems(filter interface{}, destination interface{}) (error) {
	cur, err := t.Collection().Find(context.Background(), filter)

	if err != nil {
		logger.Fatal(err)
	}

	dstv := reflect.ValueOf(destination)
	//itemV := dstv.Elem()
	logger.Warnf("DEST KIND: %v %v", dstv.Elem().Kind() == reflect.Slice, dstv.Elem().Kind() == reflect.Map)

	//if dstv.IsNil() || dstv.Kind() != reflect.Ptr {
	//	return errors.New("emit macho dwarf: elf header corrupted")
	//}

	if !(dstv.Elem().Kind() == reflect.Slice || dstv.Elem().Kind() == reflect.Map) {
		return ErrExpectingSlicePointer // todo better error msg
	}

	if dstv.Kind() != reflect.Ptr || dstv.IsNil() {
		return ErrExpectingSliceMapStruct
	}

	slicev := dstv.Elem()
	itemT := slicev.Type().Elem()

	reset(destination)

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		raw, err := cur.DecodeBytes()
		if err != nil {
			return err
		}

		//////
		var item reflect.Value

		switch itemT.Kind() {
		case reflect.Map:
			item = reflect.MakeMap(itemT)
		case reflect.Struct:
			fmt.Printf("making new struct\n")
			item = reflect.New(itemT)
		case reflect.Ptr:
			objT := itemT.Elem()
			if objT.Kind() != reflect.Struct {
				// todo error
				//return item, ErrExpectingMapOrStruct
				fmt.Printf("ERRRORRRR \n")
			}
			item = reflect.New(objT)
		}

		err = inferFields(item, raw)
		if err != nil {
			// todo
		}
		/// dest
		if dstv.Elem().Kind() == reflect.Slice {
			if itemT.Kind() == reflect.Ptr {
				slicev = reflect.Append(slicev, item)
			} else {
				slicev = reflect.Append(slicev, reflect.Indirect(item))
			}
		} else {
			if obj, ok := item.Interface().(CollectionObject); ok {
				key := obj.GetKey()
				slicev.SetMapIndex(reflect.ValueOf(key), item.Elem())
			} else {
				panic("AHH")
			}
		}
		/////
	}
	if err := cur.Err(); err != nil {
		log.Printf("ERROR")
	}

	dstv.Elem().Set(slicev)

	return nil
}

func parseTag(raw bson.Raw, tag string) (*bson.RawValue, error) {
	tags := strings.Split(tag, ",")

	for i := 0; i < len(tags); i++ {
		val, err := raw.LookupErr(tags[i])
		if err == nil {
			return &val, nil
		}
	}

	return nil, ErrMissingField
}

func inferFields(dest reflect.Value, raw bson.Raw) (error) {
	for i := 0; i < dest.Elem().NumField(); i++ {
		tag := dest.Elem().Type().Field(i).Tag.Get("bson")
		field := dest.Elem().Field(i)

		fmt.Printf("%d: %s `%v` %s = %v\n", i,
			dest.Elem().Type().Field(i).Name, tag, field.Type().String(), field.Interface())

		if field.CanSet() {
			var rv reflect.Value
			val, err := parseTag(raw, tag)

			if err != nil {
				// todo how to handle new fields? they will produce error here
				logger.Errorf("LOOKUP ERR %v", err)
				//return err
			}

			// null entries will fail here use xxxOK versions of method
			// TODO all types?
			fts := field.Type().String() // something something dynamic types
			switch fts {
			case "string":
				rv = reflect.ValueOf(val.String())
			case "objectid.ObjectID":
				v := raw.Lookup("_id")
				val = &v //raw.Lookup("_id")
				rv = reflect.ValueOf(val.ObjectID())
			case "time.Time":
				time, _ := val.TimeOK()
				rv = reflect.ValueOf(time)
			default:
				return fmt.Errorf("UNSUPPORTED TYPE %v", fts)
			}

			fmt.Printf("     rv %v val %v\n", rv, val)
			field.Set(rv)
		}
	}

	return nil
}

func reset(data interface{}) error {
	// Resetting element.
	v := reflect.ValueOf(data).Elem()
	t := v.Type()

	var z reflect.Value

	switch v.Kind() {
	case reflect.Map:
		z = reflect.MakeMap(t)
	case reflect.Slice:
		z = reflect.MakeSlice(t, 0, v.Cap())
	default:
		z = reflect.Zero(t)
	}

	v.Set(z)
	return nil
}
