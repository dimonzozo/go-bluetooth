package service

import (
	"github.com/godbus/dbus/v5"
	log "github.com/sirupsen/logrus"
)

//ReadValue read a value
func (s *Char) ReadValue(options map[string]interface{}) ([]byte, *dbus.Error) {

	log.Debug("Characteristic.ReadValue")
	if s.readCallback != nil {
		b, err := s.readCallback(s, options)
		if err != nil {
			return nil, dbus.MakeFailedError(err)
		}
		return b, nil
	}

	return s.Properties.Value, nil
}

//WriteValue write a value
func (s *Char) WriteValue(value []byte, options map[string]interface{}) *dbus.Error {

	log.Trace("Characteristic.WriteValue")

	val := value
	if s.writeCallback != nil {
		log.Trace("Used write callback")
		b, err := s.writeCallback(s, value)
		val = b
		if err != nil {
			return dbus.MakeFailedError(err)
		}
	} else {
		log.Trace("Store directly to value (no callback)")
	}

	// TODO update on Properties interface
	s.Properties.Value = val
	s.iprops.Instance().Set(s.Interface(), "Value", dbus.MakeVariant(value))

	return nil
}
