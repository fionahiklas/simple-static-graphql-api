package teststorage

import (
	"github.com/fionahiklas/simple-static-graphql-api/pkg/alarmstorage"
)

var sensors = map[string]*alarmstorage.Sensor{
	"sensor_01": {
		Id:          "sensor_01",
		Name:        "Sensor One",
		Description: "",
	},

	"sensor_02": {
		Id:          "sensor_02",
		Name:        "Sensor Two",
		Description: "",
	},

	"sensor_03": {
		Id:          "sensor_03",
		Name:        "Sensor Three",
		Description: "",
	},

	"sensor_04": {
		Id:          "sensor_04",
		Name:        "Sensor Four",
		Description: "",
	},

	"sensor_05": {
		Id:          "sensor_05",
		Name:        "Sensor Five",
		Description: "",
	},

	"sensor_06": {
		Id:          "sensor_06",
		Name:        "Sensor Six",
		Description: "",
	},
}

var alarms = map[string]*alarmstorage.Alarm{
	"alarm_01": {
		Id:          "alarm_01",
		Name:        "Alarm One",
		Description: "",
		Sensors:     []*alarmstorage.Sensor{sensors["sensor_01"], sensors["sensor_02"]},
	},

	"alarm_02": {
		Id:          "alarm_02",
		Name:        "Alarm Two",
		Description: "",
		Sensors:     []*alarmstorage.Sensor{sensors["sensor_03"], sensors["sensor_04"]},
	},

	"alarm_03": {
		Id:          "alarm_03",
		Name:        "Alarm Three",
		Description: "",
		Sensors:     []*alarmstorage.Sensor{sensors["sensor_05"], sensors["sensor_06"]},
	},
}

var homes = map[string]*alarmstorage.Home{
	"home_01": {
		Id:          "home_01",
		Name:        "Home One",
		Description: "",
		Alarm:       alarms["alarm_01"],
	},

	"home_02": {
		Id:          "home_02",
		Name:        "Home Two",
		Description: "",
		Alarm:       alarms["alarm_02"],
	},
}
