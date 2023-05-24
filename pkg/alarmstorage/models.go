package alarmstorage

type Sensor struct {
	Identifier  string
	Name        string
	Description string
}

type Alarm struct {
	Identifier  string
	Name        string
	Description string
	Sensors     []*Sensor
}

type Home struct {
	Identifier  string
	Name        string
	Description string
	Alarm       *Alarm
}
