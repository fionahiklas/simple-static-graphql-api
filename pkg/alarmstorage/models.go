package alarmstorage

type Sensor struct {
	Id          string
	Name        string
	Description string
}

type Alarm struct {
	Id          string
	Name        string
	Description string
	Sensors     []*Sensor
}

type Home struct {
	Id          string
	Name        string
	Description string
	Alarm       *Alarm
}
