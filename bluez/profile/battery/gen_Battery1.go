package battery



import (
	"github.com/godbus/dbus/v5"
	"github.com/dimonzozo/go-bluetooth/bluez"
	"github.com/dimonzozo/go-bluetooth/props"
	"github.com/dimonzozo/go-bluetooth/util"
	"sync"
)

var Battery1Interface = "org.bluez.Battery1"


// NewBattery1 create a new instance of Battery1
//
// Args:
// - objectPath: [variable prefix]/{hci0,hci1,...}/dev_XX_XX_XX_XX_XX_XX
func NewBattery1(objectPath dbus.ObjectPath) (*Battery1, error) {
	a := new(Battery1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: Battery1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(Battery1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
Battery1 Battery hierarchy

*/
type Battery1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*Battery1Properties
}

// Battery1Properties contains the exposed properties of an interface
type Battery1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
	Percentage The percentage of battery left as an unsigned 8-bit integer.
	*/
	Percentage byte

}

//Lock access to properties
func (p *Battery1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *Battery1Properties) Unlock() {
	p.lock.Unlock()
}






// GetPercentage get Percentage value
func (a *Battery1) GetPercentage() (byte, error) {
	v, err := a.GetProperty("Percentage")
	if err != nil {
		return byte(0), err
	}
	return v.Value().(byte), nil
}



// Close the connection
func (a *Battery1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return Battery1 object path
func (a *Battery1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return Battery1 dbus client
func (a *Battery1) Client() *bluez.Client {
	return a.client
}

// Interface return Battery1 interface
func (a *Battery1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *Battery1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}


// ToMap convert a Battery1Properties to map
func (a *Battery1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an Battery1Properties
func (a *Battery1Properties) FromMap(props map[string]interface{}) (*Battery1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an Battery1Properties
func (a *Battery1Properties) FromDBusMap(props map[string]dbus.Variant) (*Battery1Properties, error) {
	s := new(Battery1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *Battery1) ToProps() bluez.Properties {
	return a.Properties
}

// GetProperties load all available properties
func (a *Battery1) GetProperties() (*Battery1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *Battery1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *Battery1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *Battery1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *Battery1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *Battery1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *Battery1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}




