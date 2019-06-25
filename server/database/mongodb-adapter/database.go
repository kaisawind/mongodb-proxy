package mongodbadapter

import (
	"gopkg.in/mgo.v2"
)

// DatabaseNames get all database names
func (a *adapter) DatabaseNames() (names []string, err error) {
	return a.session.DatabaseNames()
}

// CollectionNames get database's collection names
func (a *adapter) CollectionNames(dbName string) (names []string, err error) {
	return a.session.DB(dbName).CollectionNames()
}

func (a *adapter) Collection(dbName, cName string) *mgo.Collection {
	return a.session.DB(dbName).C(cName)
}

func (a *adapter) BuildInfo() (mgo.BuildInfo, error) {
	return a.session.BuildInfo()
}

func (a *adapter) LiveServers() (addrs []string) {
	return a.session.LiveServers()
}
