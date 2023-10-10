package mongodb

import (
	"context"

	"github.com/debeando/go-common/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Context context.Context
	Client  *mongo.Client
	Name    string
	DSN     string
}

var instance = make(map[string]*Connection)

func New(name, dsn string) *Connection {
	if instance[name] == nil {
		instance[name] = &Connection{
			Name: name,
			DSN:  dsn,
		}
		instance[name].Name = name
	}
	return instance[name]
}

func (c *Connection) Connect() (err error) {
	clientOptions := options.Client().ApplyURI(c.DSN)
	c.Client, err = mongo.Connect(c.Context, clientOptions)
	if err != nil {
		log.ErrorWithFields("MongoDB", log.Fields{"message": err})
		return err
	}

	err = c.Client.Ping(c.Context, nil)
	if err != nil {
		log.ErrorWithFields("MongoDB", log.Fields{"message": err})
		return err
	}

	return nil
}

func (c *Connection) GetServerStatus() (ss map[string]interface{}) {
	db := c.Client.Database("admin")
	cmd := bson.D{primitive.E{Key: "serverStatus", Value: "1"}}
	db.RunCommand(context.TODO(), cmd).Decode(&ss)

	return ss
}
