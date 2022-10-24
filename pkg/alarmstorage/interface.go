package alarmstorage

type ReadFromStorage interface {
	GetAlarm(id string) *Alarm
	GetAlarms() []*Alarm
	GetHome(id string) *Home
	GetHomes() []*Home
	GetSensor(id string) *Sensor
	GetSensors() []*Sensor
}

type WriteToStorage interface {
	CreateHome(home *Home)
}

type ReadAndWrite interface {
	ReadFromStorage
	WriteToStorage
}
