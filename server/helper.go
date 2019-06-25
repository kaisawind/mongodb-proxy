package server

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/kaisawind/mongodb-proxy/server/database"
	dbadapter "github.com/kaisawind/mongodb-proxy/server/database/mongodb-adapter"
	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Default value
const (
	EnvPrefix           = "proxy"
	MongoAddressKey     = "mongo.address"
	MongoAddressDefault = "localhost:27017"
)

// database url
const (
	ProductURL         = "products/products"
	AertlistURL        = "alert/alertlist"
	AlarmlistURL       = "alert/alarmlist"
	DeviceURL          = "devices/"
	LogrdbURL          = "logrdb/"
	UsersURL           = "users/users"
	CubeURL            = "cube/cube_message"
	CubeVoltageURL     = "cube/cube_message/voltage"
	CubeTemperatureURL = "cube/cube_message/temperature"
	CubeCurrentURL     = "cube/cube_message/current"
	CubeFrequencyURL   = "cube/cube_message/frequency"
)

var (
	// Adapter mongodb instance
	Adapter database.Adapter
	once    sync.Once
)

func init() {
	logrus.Infoln("initViper")
	initViper()
	Adapter = dbadapter.NewAdapter(MongoAddress())
	info, err := Adapter.BuildInfo()
	if err != nil {
		logrus.Errorln("BuildInfo Error:", err)
		panic(err)
	}
	logrus.Infoln("Version:", info.Version)
	logrus.Infoln("SysInfo:", info.SysInfo)
	logrus.Infoln("MaxObjectSize:", info.MaxObjectSize)

	addrs := Adapter.LiveServers()
	for _, addr := range addrs {
		logrus.Infoln("addr:", addr)
	}
}

// InitViper ...
func initViper() {
	once.Do(func() {
		viper.SetEnvPrefix(EnvPrefix)
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

		viper.SetDefault(MongoAddressKey, MongoAddressDefault)
	})
}

// MongoAddress get mongo address from env(default:"localhost:27017")
func MongoAddress() string {
	val := viper.GetString(MongoAddressKey)
	logrus.Infoln("MongoAddress:", val)
	if val == "" {
		val = MongoAddressDefault
	}
	return val
}

// Error ...
func Error(err error) middleware.Responder {
	logrus.Errorln("Returningerror:", err)
	return middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
		code := http.StatusInternalServerError
		w.WriteHeader(code)
		payload, _ := json.Marshal(map[string]interface{}{
			"error": err.Error(),
			"code":  code,
		})
		w.Write(payload)
	})
}

// Timeserie ...
type Timeserie struct {
	Total     int       `json:"total,omitempty" bson:"total,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}

// Cube ...
type Cube struct {
	Voltage     int   `json:"voltage,omitempty" bson:"voltage,omitempty"`
	Temperature int   `json:"temperature,omitempty" bson:"temperature,omitempty"`
	Current     int   `json:"current,omitempty" bson:"current,omitempty"`
	Frequency   int   `json:"frequency,omitempty" bson:"frequency,omitempty"`
	Timestamp   int64 `json:"timeStamp,omitempty" bson:"timeStamp,omitempty"`
}
