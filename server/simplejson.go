package server

import (
	"strings"
	"time"

	"github.com/kaisawind/mongodb-proxy/restapi/operations/simple_json"
	"github.com/kaisawind/mongodb-proxy/v1"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

// TestDatasource GET /
func TestDatasource(params simple_json.TestDatasourceParams) middleware.Responder {
	logrus.Infoln("TestDatasourceParams:", params)
	return simple_json.NewTestDatasourceOK()
}

// AnnotationQuery POST /annotations
func AnnotationQuery(params simple_json.AnnotationQueryParams) middleware.Responder {
	logrus.Infoln("AnnotationQueryParams:", params.Options)
	payload := &v1.Annotations{}
	return simple_json.NewAnnotationQueryOK().WithPayload(payload)
}

// MetricFindQuery POST /search
func MetricFindQuery(params simple_json.MetricFindQueryParams) middleware.Responder {
	logrus.Infoln("MetricFindQueryParams:", params.Options)
	dbnames, err := Adapter.DatabaseNames()
	if err != nil {
		return Error(err)
	}
	payload := []string{}

	for _, dbname := range dbnames {
		cnames, err := Adapter.CollectionNames(dbname)
		if err != nil {
			logrus.Errorln("CollectionNames Error:", err)
			continue
		}
		for _, cname := range cnames {
			name := dbname + "/" + cname
			payload = append(payload, name)
		}
	}
	return simple_json.NewMetricFindQueryOK().WithPayload(payload)
}

// Query POST /query
func Query(params simple_json.QueryParams) middleware.Responder {
	logrus.Infoln("QueryParams:", params.Options)
	targets := params.Options.Targets

	payload := v1.Timeseries{}

	for _, target := range targets {
		names := strings.Split(target.Target, "/")
		logrus.Infoln("Query target:", names)
		if len(names) < 2 {
			continue
		}
		if target.Type != nil && *target.Type == "timeserie" {
			timeserie := &v1.Timeserie{
				Datapoints: [][]interface{}{},
				Target:     names[1],
			}
			if target.Target == CubeVoltageURL {
				series, err := GetCubeVoltageTimeserie(names[0], names[1], params.Options)
				if err != nil {
					continue
				}
				for _, serie := range series {
					logrus.Debugln(CubeURL, *serie)
					point := []interface{}{serie.Total, serie.Timestamp.Unix() * 1000}
					timeserie.Datapoints = append(timeserie.Datapoints, point)
				}
				payload = append(payload, timeserie)
			}
			if target.Target == CubeTemperatureURL {
				series, err := GetCubeTemperatureTimeserie(names[0], names[1], params.Options)
				if err != nil {
					continue
				}
				for _, serie := range series {
					logrus.Debugln(CubeURL, *serie)
					point := []interface{}{serie.Total, serie.Timestamp.Unix() * 1000}
					timeserie.Datapoints = append(timeserie.Datapoints, point)
				}
				payload = append(payload, timeserie)
			}
			if target.Target == CubeCurrentURL {
				series, err := GetCubeCurrentTimeserie(names[0], names[1], params.Options)
				if err != nil {
					continue
				}
				for _, serie := range series {
					logrus.Debugln(CubeURL, *serie)
					point := []interface{}{serie.Total, serie.Timestamp.Unix() * 1000}
					timeserie.Datapoints = append(timeserie.Datapoints, point)
				}
				payload = append(payload, timeserie)
			}
			if target.Target == CubeFrequencyURL {
				series, err := GetCubeFrequencyTimeserie(names[0], names[1], params.Options)
				if err != nil {
					continue
				}
				for _, serie := range series {
					logrus.Debugln(CubeURL, *serie)
					point := []interface{}{serie.Total, serie.Timestamp.Unix() * 1000}
					timeserie.Datapoints = append(timeserie.Datapoints, point)
				}
				payload = append(payload, timeserie)
			}
			if target.Target == CubeURL {
				series, err := GetCubeTimeserie(names[0], names[1], params.Options)
				if err != nil {
					continue
				}
				for _, serie := range series {
					logrus.Debugln(CubeURL, *serie)
					point := []interface{}{serie.Total, serie.Timestamp.Unix() * 1000}
					timeserie.Datapoints = append(timeserie.Datapoints, point)
				}
				payload = append(payload, timeserie)
			}
			if target.Target == UsersURL {
				series, err := GetUsersTimeserie(names[0], names[1], params.Options)
				if err != nil {
					continue
				}
				for _, serie := range series {
					point := []interface{}{serie.Total, serie.Timestamp.Unix() * 1000}
					timeserie.Datapoints = append(timeserie.Datapoints, point)
				}
				payload = append(payload, timeserie)
			}
			if target.Target == ProductURL {
				series, err := GetProductsTimeserie(names[0], names[1], params.Options)
				if err != nil {
					continue
				}
				for _, serie := range series {
					point := []interface{}{serie.Total, serie.Timestamp.Unix() * 1000}
					timeserie.Datapoints = append(timeserie.Datapoints, point)
				}
				payload = append(payload, timeserie)
			}
			if strings.HasPrefix(target.Target, DeviceURL) {
				if target.Target == DeviceURL {
					cnames, err := Adapter.CollectionNames(names[0])
					if err != nil {
						continue
					}
					for _, cname := range cnames {
						timeserie = &v1.Timeserie{
							Datapoints: [][]interface{}{},
							Target:     cname,
						}
						series, err := GetDevicesTimeserie(names[0], cname, params.Options)
						if err != nil {
							continue
						}
						for _, serie := range series {
							point := []interface{}{serie.Total, serie.Timestamp.Unix() * 1000}
							timeserie.Datapoints = append(timeserie.Datapoints, point)
						}
						payload = append(payload, timeserie)
					}
				} else {
					series, err := GetDevicesTimeserie(names[0], names[1], params.Options)
					if err != nil {
						continue
					}
					for _, serie := range series {
						point := []interface{}{serie.Total, serie.Timestamp.Unix() * 1000}
						timeserie.Datapoints = append(timeserie.Datapoints, point)
					}
					payload = append(payload, timeserie)
				}
			}
			if strings.HasPrefix(target.Target, LogrdbURL) {
				if target.Target == LogrdbURL {
					cnames, err := Adapter.CollectionNames(names[0])
					if err != nil {
						continue
					}
					for _, cname := range cnames {
						timeserie = &v1.Timeserie{
							Datapoints: [][]interface{}{},
							Target:     cname,
						}
						series, err := GetLogrdbTimeserie(names[0], cname, params.Options)
						if err != nil {
							continue
						}
						for _, serie := range series {
							point := []interface{}{serie.Total, serie.Timestamp.Unix() * 1000}
							timeserie.Datapoints = append(timeserie.Datapoints, point)
						}
						payload = append(payload, timeserie)
					}
				} else {
					series, err := GetLogrdbTimeserie(names[0], names[1], params.Options)
					if err != nil {
						continue
					}
					for _, serie := range series {
						point := []interface{}{serie.Total, serie.Timestamp.Unix() * 1000}
						timeserie.Datapoints = append(timeserie.Datapoints, point)
					}
					payload = append(payload, timeserie)
				}
			}
		}
	}

	return simple_json.NewQueryOK().WithPayload(payload)
}

// GetProductsTimeserie ...
/*
db.getCollection('products').aggregate([
	{
		'$match':{
			'created_at': {'$gte': "2019-03-13T10:30:56Z", '$lte': "2019-03-14T09:25:50Z"}
		}
	},
	{
		'$group': {
			'_id': {
				'$subtract': [
					{'$toDate': '$created_at'},
					{
						'$mod': [
							{'$toLong': {'$toDate': '$created_at'}},
							60 * 60 * 24 * 1000
						]
					}
				]
			},
			'total': {'$sum': 1}
		}
	},
	{
		'$project': { '_id': 0, 'timestamp': '$_id', 'total': 1}
	},
	{
		'$sort': { 'timestamp': 1 }
	}
])
*/
func GetProductsTimeserie(dbName, cName string, options *v1.Query) ([]*Timeserie, error) {
	timeseries := []*Timeserie{}

	timerange := bson.M{"$gte": options.Range.From.String(), "$lte": options.Range.To.String()}
	match := bson.M{"$match": bson.M{"created_at": timerange}}
	mod := bson.M{"$mod": []interface{}{bson.M{"$toLong": bson.M{"$toDate": "$created_at"}}, options.IntervalMs}}
	subtract := bson.M{"$subtract": []interface{}{bson.M{"$toDate": "$created_at"}, mod}}
	group := bson.M{
		"$group": bson.M{
			"_id":   subtract,
			"total": bson.M{"$sum": 1},
		},
	}
	project := bson.M{"$project": bson.M{
		"timestamp": "$_id",
		"_id":       0,
		"total":     1,
	}}
	sort := bson.M{"$sort": bson.M{"timestamp": 1}}
	pipeline := []bson.M{match, group, project, sort}
	logrus.Debugln("GetProductsTimeserie pipeline:", pipeline)
	err := Adapter.Collection(dbName, cName).Pipe(pipeline).All(&timeseries)
	if err != nil {
		logrus.Errorln("GetProductsTimeserie Error:", err)
		return timeseries, err
	}
	logrus.Debugln("GetProductsTimeserie timeseries:", timeseries)
	return timeseries, nil
}

// GetDevicesTimeserie ...
func GetDevicesTimeserie(dbName, cName string, options *v1.Query) ([]*Timeserie, error) {
	timeseries := []*Timeserie{}

	timerange := bson.M{"$gte": options.Range.From.String(), "$lte": options.Range.To.String()}
	match := bson.M{"$match": bson.M{"created_at": timerange}}
	mod := bson.M{"$mod": []interface{}{bson.M{"$toLong": bson.M{"$toDate": "$created_at"}}, options.IntervalMs}}
	subtract := bson.M{"$subtract": []interface{}{bson.M{"$toDate": "$created_at"}, mod}}
	group := bson.M{
		"$group": bson.M{
			"_id":   subtract,
			"total": bson.M{"$sum": 1},
		},
	}
	project := bson.M{"$project": bson.M{
		"timestamp": "$_id",
		"_id":       0,
		"total":     1,
	}}
	sort := bson.M{"$sort": bson.M{"timestamp": 1}}
	pipeline := []bson.M{match, group, project, sort}
	logrus.Debugln("GetDevicesTimeserie pipeline:", pipeline)
	err := Adapter.Collection(dbName, cName).Pipe(pipeline).All(&timeseries)
	if err != nil {
		logrus.Errorln("GetDevicesTimeserie Error", err)
		return timeseries, err
	}
	logrus.Debugln("GetDevicesTimeserie timeseries:", timeseries)
	return timeseries, nil
}

// GetLogrdbTimeserie ...
func GetLogrdbTimeserie(dbName, cName string, options *v1.Query) ([]*Timeserie, error) {
	timeseries := []*Timeserie{}

	timerange := bson.M{"$gte": time.Time(options.Range.From).Unix(), "$lte": time.Time(options.Range.To).Unix()}
	match := bson.M{"$match": bson.M{"meta.time": timerange}}
	mod := bson.M{"$mod": []interface{}{"$meta.time", options.IntervalMs}}
	subtract := bson.M{"$subtract": []interface{}{"$meta.time", mod}}
	group := bson.M{
		"$group": bson.M{
			"_id":   subtract,
			"total": bson.M{"$sum": 1},
		},
	}
	project := bson.M{"$project": bson.M{
		"timestamp": bson.M{
			"$toDate": bson.M{
				"$multiply": []interface{}{
					"$_id",
					1000,
				},
			},
		},
		"_id":   0,
		"total": 1,
	}}
	sort := bson.M{"$sort": bson.M{"timestamp": 1}}
	pipeline := []bson.M{match, group, project, sort}
	logrus.Debugln("GetLogrdbTimeserie pipeline:", pipeline)
	err := Adapter.Collection(dbName, cName).Pipe(pipeline).All(&timeseries)
	if err != nil {
		logrus.Errorln("GetLogrdbTimeserie Error", err)
		return timeseries, err
	}
	logrus.Debugln("GetLogrdbTimeserie timeseries:", timeseries)
	return timeseries, nil
}

// GetUsersTimeserie ...
func GetUsersTimeserie(dbName, cName string, options *v1.Query) ([]*Timeserie, error) {
	timeseries := []*Timeserie{}

	timerange := bson.M{"$gte": options.Range.From.String(), "$lte": options.Range.To.String()}
	match := bson.M{"$match": bson.M{"created_at": timerange}}
	mod := bson.M{"$mod": []interface{}{bson.M{"$toLong": bson.M{"$toDate": "$created_at"}}, options.IntervalMs}}
	subtract := bson.M{"$subtract": []interface{}{bson.M{"$toDate": "$created_at"}, mod}}
	group := bson.M{
		"$group": bson.M{
			"_id":   subtract,
			"total": bson.M{"$sum": 1},
		},
	}
	project := bson.M{"$project": bson.M{
		"timestamp": "$_id",
		"_id":       0,
		"total":     1,
	}}
	sort := bson.M{"$sort": bson.M{"timestamp": 1}}
	pipeline := []bson.M{match, group, project, sort}
	logrus.Debugln("GetUsersTimeserie pipeline:", pipeline)
	err := Adapter.Collection(dbName, cName).Pipe(pipeline).All(&timeseries)
	if err != nil {
		logrus.Errorln("GetUsersTimeserie Error:", err)
		return timeseries, err
	}
	logrus.Debugln("GetUsersTimeserie timeseries:", timeseries)
	return timeseries, nil
}

// GetCubeTimeserie ...
func GetCubeTimeserie(dbName, cName string, options *v1.Query) ([]*Timeserie, error) {
	timeseries := []*Timeserie{}

	timerange := bson.M{"$gte": time.Time(options.Range.From).Unix(), "$lte": time.Time(options.Range.To).Unix()}
	match := bson.M{"$match": bson.M{"timeStamp": timerange}}
	mod := bson.M{"$mod": []interface{}{"$timeStamp", options.IntervalMs}}
	subtract := bson.M{"$subtract": []interface{}{"$timeStamp", mod}}
	group := bson.M{
		"$group": bson.M{
			"_id":   subtract,
			"total": bson.M{"$sum": 1},
		},
	}
	project := bson.M{"$project": bson.M{
		"timestamp": bson.M{
			"$toDate": bson.M{
				"$multiply": []interface{}{
					"$_id",
					1000,
				},
			},
		},
		"_id":   0,
		"total": 1,
	}}
	sort := bson.M{"$sort": bson.M{"timestamp": 1}}
	pipeline := []bson.M{match, group, project, sort}
	logrus.Debugln("GetCubeTimeserie pipeline:", pipeline)
	err := Adapter.Collection(dbName, cName).Pipe(pipeline).All(&timeseries)
	if err != nil {
		logrus.Errorln("GetCubeTimeserie Error", err)
		return timeseries, err
	}
	logrus.Debugln("GetCubeTimeserie timeseries:", timeseries)
	return timeseries, nil
}

// GetCubeVoltageTimeserie ...
func GetCubeVoltageTimeserie(dbName, cName string, options *v1.Query) ([]*Timeserie, error) {
	timeseries := []*Timeserie{}

	timerange := bson.M{"$gte": time.Time(options.Range.From).Unix(), "$lte": time.Time(options.Range.To).Unix()}
	cubes := []*Cube{}
	err := Adapter.Collection(dbName, cName).Find(bson.M{"timeStamp": timerange}).Limit(int(options.MaxDataPoints)).All(&cubes)
	if err != nil {
		logrus.Errorln("GetCubeVoltageTimeserie Error", err)
		return timeseries, err
	}
	for _, cube := range cubes {
		timeserie := &Timeserie{
			Total:     cube.Voltage,
			Timestamp: time.Unix(cube.Timestamp, 0),
		}
		timeseries = append(timeseries, timeserie)
	}
	logrus.Debugln("GetCubeVoltageTimeserie timeseries:", timeseries)
	return timeseries, nil
}

// GetCubeTemperatureTimeserie ...
func GetCubeTemperatureTimeserie(dbName, cName string, options *v1.Query) ([]*Timeserie, error) {
	timeseries := []*Timeserie{}

	timerange := bson.M{"$gte": time.Time(options.Range.From).Unix(), "$lte": time.Time(options.Range.To).Unix()}
	cubes := []*Cube{}
	err := Adapter.Collection(dbName, cName).Find(bson.M{"timeStamp": timerange}).Limit(int(options.MaxDataPoints)).All(&cubes)
	if err != nil {
		logrus.Errorln("GetCubeTemperatureTimeserie Error", err)
		return timeseries, err
	}
	for _, cube := range cubes {
		timeserie := &Timeserie{
			Total:     cube.Temperature,
			Timestamp: time.Unix(cube.Timestamp, 0),
		}
		timeseries = append(timeseries, timeserie)
	}
	logrus.Debugln("GetCubeTemperatureTimeserie timeseries:", timeseries)
	return timeseries, nil
}

// GetCubeCurrentTimeserie ...
func GetCubeCurrentTimeserie(dbName, cName string, options *v1.Query) ([]*Timeserie, error) {
	timeseries := []*Timeserie{}

	timerange := bson.M{"$gte": time.Time(options.Range.From).Unix(), "$lte": time.Time(options.Range.To).Unix()}
	cubes := []*Cube{}
	err := Adapter.Collection(dbName, cName).Find(bson.M{"timeStamp": timerange}).Limit(int(options.MaxDataPoints)).All(&cubes)
	if err != nil {
		logrus.Errorln("GetCubeCurrentTimeserie Error", err)
		return timeseries, err
	}
	for _, cube := range cubes {
		timeserie := &Timeserie{
			Total:     cube.Current,
			Timestamp: time.Unix(cube.Timestamp, 0),
		}
		timeseries = append(timeseries, timeserie)
	}
	logrus.Debugln("GetCubeCurrentTimeserie timeseries:", timeseries)
	return timeseries, nil
}

// GetCubeFrequencyTimeserie ...
func GetCubeFrequencyTimeserie(dbName, cName string, options *v1.Query) ([]*Timeserie, error) {
	timeseries := []*Timeserie{}

	timerange := bson.M{"$gte": time.Time(options.Range.From).Unix(), "$lte": time.Time(options.Range.To).Unix()}
	cubes := []*Cube{}
	err := Adapter.Collection(dbName, cName).Find(bson.M{"timeStamp": timerange}).Limit(int(options.MaxDataPoints)).All(&cubes)
	if err != nil {
		logrus.Errorln("GetCubeFrequencyeTimeserie Error", err)
		return timeseries, err
	}
	for _, cube := range cubes {
		timeserie := &Timeserie{
			Total:     cube.Frequency,
			Timestamp: time.Unix(cube.Timestamp, 0),
		}
		timeseries = append(timeseries, timeserie)
	}
	logrus.Debugln("GetCubeFrequencyeTimeserie timeseries:", timeseries)
	return timeseries, nil
}
