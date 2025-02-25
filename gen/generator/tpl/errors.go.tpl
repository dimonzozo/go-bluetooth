
package profile

import (
	"github.com/godbus/dbus/v5"
)

var (
{{ range .List }}
	// {{.Name}} map to org.bluez.Error.{{.Name}}
	Err{{.Name}} = dbus.Error{
		Name: "org.bluez.Error.{{.Name}}",
		Body: []interface{}{"{{.Name}}"},
	}
{{ end }}
)
