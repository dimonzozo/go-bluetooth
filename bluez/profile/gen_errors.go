
package profile

import (
	"github.com/godbus/dbus/v5"
)

var (

	// NotReady map to org.bluez.Error.NotReady
	ErrNotReady = dbus.Error{
		Name: "org.bluez.Error.NotReady",
		Body: []interface{}{"NotReady"},
	}

	// InvalidArguments map to org.bluez.Error.InvalidArguments
	ErrInvalidArguments = dbus.Error{
		Name: "org.bluez.Error.InvalidArguments",
		Body: []interface{}{"InvalidArguments"},
	}

	// Failed map to org.bluez.Error.Failed
	ErrFailed = dbus.Error{
		Name: "org.bluez.Error.Failed",
		Body: []interface{}{"Failed"},
	}

	// NotPermitted map to org.bluez.Error.NotPermitted
	ErrNotPermitted = dbus.Error{
		Name: "org.bluez.Error.NotPermitted",
		Body: []interface{}{"NotPermitted"},
	}

	// DoesNotExist map to org.bluez.Error.DoesNotExist
	ErrDoesNotExist = dbus.Error{
		Name: "org.bluez.Error.DoesNotExist",
		Body: []interface{}{"DoesNotExist"},
	}

	// Rejected map to org.bluez.Error.Rejected
	ErrRejected = dbus.Error{
		Name: "org.bluez.Error.Rejected",
		Body: []interface{}{"Rejected"},
	}

	// NotConnected map to org.bluez.Error.NotConnected
	ErrNotConnected = dbus.Error{
		Name: "org.bluez.Error.NotConnected",
		Body: []interface{}{"NotConnected"},
	}

	// NotAcquired map to org.bluez.Error.NotAcquired
	ErrNotAcquired = dbus.Error{
		Name: "org.bluez.Error.NotAcquired",
		Body: []interface{}{"NotAcquired"},
	}

	// NotSupported map to org.bluez.Error.NotSupported
	ErrNotSupported = dbus.Error{
		Name: "org.bluez.Error.NotSupported",
		Body: []interface{}{"NotSupported"},
	}

	// NotAuthorized map to org.bluez.Error.NotAuthorized
	ErrNotAuthorized = dbus.Error{
		Name: "org.bluez.Error.NotAuthorized",
		Body: []interface{}{"NotAuthorized"},
	}

	// NotAvailable map to org.bluez.Error.NotAvailable
	ErrNotAvailable = dbus.Error{
		Name: "org.bluez.Error.NotAvailable",
		Body: []interface{}{"NotAvailable"},
	}

)
