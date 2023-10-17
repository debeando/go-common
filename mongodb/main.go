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
		log.ErrorWithFields("MongoDB:Connect", log.Fields{"message": err})
		return err
	}

	err = c.Client.Ping(c.Context, nil)
	if err != nil {
		log.ErrorWithFields("MongoDB:Connect:Ping", log.Fields{"message": err})
		return err
	}

	return nil
}

func (c *Connection) Close() {
	if err := c.Client.Disconnect(c.Context); err != nil {
		log.ErrorWithFields("MongoDB:Close", log.Fields{"message": err})
	}
}

func (c *Connection) RunCommand(database string, cmd interface{}, result interface{}) error {
	db := c.Client.Database(database)
	return db.RunCommand(context.TODO(), cmd).Decode(result)
}

func (c *Connection) ServerStatus() *ServerStatus {
	serverstatus := &ServerStatus{}
	cmd := bson.D{primitive.E{Key: "serverStatus", Value: "1"}}
	c.RunCommand("admin", cmd, serverstatus)

	return serverstatus
}

func (c *Connection) Databases() *Databases {
	databases := &Databases{}
	cmd := bson.D{primitive.E{Key: "listDatabases", Value: "1"}}
	c.RunCommand("admin", cmd, databases)

	return databases
}

func (c *Connection) Collections(database string) []string {
	colls, err := c.Client.Database(database).ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		log.ErrorWithFields("MongoDB:Collections", log.Fields{"message": err})
	}

	return colls
}

func (c *Connection) CollectionStats(dbName, colName string) *CollectionStats {
	collectionStats := &CollectionStats{}
	cmd := bson.D{primitive.E{Key: "collStats", Value: colName}}
	c.RunCommand(dbName, cmd, collectionStats)

	return collectionStats
}
