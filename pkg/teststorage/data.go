package teststorage

import (
	"github.com/fionahiklas/simple-static-graphql-api/pkg/alarmstorage"
)

var sensors = map[string]*alarmstorage.Sensor{
	"sensor_01": {
		Identifier:  "sensor_01",
		Name:        "Sensor One",
		Description: "",
	},

	"sensor_02": {
		Identifier:  "sensor_02",
		Name:        "Sensor Two",
		Description: "",
	},

	"sensor_03": {
		Identifier:  "sensor_03",
		Name:        "Sensor Three",
		Description: "",
	},

	"sensor_04": {
		Identifier:  "sensor_04",
		Name:        "Sensor Four",
		Description: "",
	},

	"sensor_05": {
		Identifier:  "sensor_05",
		Name:        "Sensor Five",
		Description: "",
	},

	"sensor_06": {
		Identifier:  "sensor_06",
		Name:        "Sensor Six",
		Description: "",
	},
}

var alarms = map[string]*alarmstorage.Alarm{
	"alarm_01": {
		Identifier:  "alarm_01",
		Name:        "Alarm One",
		Description: "",
		Sensors:     []*alarmstorage.Sensor{sensors["sensor_01"], sensors["sensor_02"]},
	},

	"alarm_02": {
		Identifier:  "alarm_02",
		Name:        "Alarm Two",
		Description: "",
		Sensors:     []*alarmstorage.Sensor{sensors["sensor_03"], sensors["sensor_04"]},
	},

	"alarm_03": {
		Identifier:  "alarm_03",
		Name:        "Alarm Three",
		Description: "",
		Sensors:     []*alarmstorage.Sensor{sensors["sensor_05"], sensors["sensor_06"]},
	},
}

var homes = map[string]*alarmstorage.Home{
	"home_01": {
		Identifier:  "home_01",
		Name:        "Home One",
		Description: "",
		Alarm:       alarms["alarm_01"],
	},

	"home_02": {
		Identifier:  "home_02",
		Name:        "Home Two",
		Description: "",
		Alarm:       alarms["alarm_02"],
	},
}
