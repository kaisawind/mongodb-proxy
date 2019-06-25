package database

import (
	"gopkg.in/mgo.v2"
)

// Adapter ...
type Adapter interface {
	LiveServers() (addrs []string)
	BuildInfo() (mgo.BuildInfo, error)
	DatabaseNames() (names []string, err error)
	CollectionNames(dbName string) (names []string, err error)
	Collection(dbName, cName string) *mgo.Collection
}
