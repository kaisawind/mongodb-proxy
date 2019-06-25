// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mongodbadapter

import (
	"runtime"

	"github.com/kaisawind/mongodb-proxy/server/database"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

const (
	dbName = "products"
)

// adapter represents the MongoDB adapter for policy storage.
type adapter struct {
	url     string
	name    string
	session *mgo.Session
}

// finalizer is the destructor for adapter.
func finalizer(a *adapter) {
	a.close()
}

// NewAdapter is the constructor for Adapter. If database name is not provided
// in the Mongo URL, 'casbin' will be used as database name.
func NewAdapter(url string) database.Adapter {
	a := &adapter{
		url:  url,
		name: "mongo",
	}

	// Open the DB, create it if not existed.
	a.open()

	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)

	return a
}

func (a *adapter) open() {
	dI, err := mgo.ParseURL(a.url)
	if err != nil {
		panic(err)
	}

	// FailFast will cause connection and query attempts to fail faster when
	// the server is unavailable, instead of retrying until the configured
	// timeout period. Note that an unavailable server may silently drop
	// packets instead of rejecting them, in which case it's impossible to
	// distinguish it from a slow server, so the timeout stays relevant.
	// dI.FailFast = true
	// Direct informs whether to establish connections only with the
	// specified seed servers, or to obtain information for the whole
	// cluster and establish connections with further servers too.
	dI.Direct = true

	if dI.Database == "" {
		dI.Database = dbName
	}

	session, err := mgo.DialWithInfo(dI)
	if err != nil {
		panic(err)
	}
	a.session = session
}

func (a *adapter) close() {
	logrus.Debugln("adapter close")
	a.session.Close()
}
