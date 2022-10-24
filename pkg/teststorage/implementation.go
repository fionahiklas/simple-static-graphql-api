package teststorage

import "github.com/fionahiklas/simple-static-graphql-api/pkg/alarmstorage"

type logger interface {
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type storage struct {
	log logger
}

func NewStorage(log logger) *storage {
	return &storage{
		log: log,
	}
}

func (st *storage) GetAlarm(id string) *alarmstorage.Alarm {
	st.log.Debugf("GetAlarm '%s' called", id)
	return alarms[id]
}

func (st *storage) GetAlarms() []*alarmstorage.Alarm {
	st.log.Debugf("GetAlarms called")
	if len(alarms) == 0 {
		return nil
	}

	resultArray := make([]*alarmstorage.Alarm, 0, len(alarms))
	for _, element := range alarms {
		resultArray = append(resultArray, element)
	}
	return resultArray
}

func (st *storage) GetHome(id string) *alarmstorage.Home {
	st.log.Debugf("GetHome '%s' called", id)
	return homes[id]
}

func (st *storage) GetHomes() []*alarmstorage.Home {
	st.log.Debugf("GetHomes called")
	if len(homes) == 0 {
		return nil
	}

	resultArray := make([]*alarmstorage.Home, 0, len(homes))
	for _, element := range homes {
		resultArray = append(resultArray, element)
	}
	return resultArray
}

func (st *storage) GetSensor(id string) *alarmstorage.Sensor {
	st.log.Debugf("GetSensor '%s' called", id)
	return sensors[id]
}

func (st *storage) GetSensors() []*alarmstorage.Sensor {
	st.log.Debugf("GetSensors called")
	if len(alarms) == 0 {
		return nil
	}

	resultArray := make([]*alarmstorage.Sensor, 0, len(alarms))
	for _, element := range sensors {
		resultArray = append(resultArray, element)
	}
	return resultArray
}

func (st *storage) CreateHome(home *alarmstorage.Home) {
	st.log.Debugf("CreateHome called")
}
