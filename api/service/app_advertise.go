package service

import (
	"github.com/dimonzozo/go-bluetooth/api"
	"github.com/dimonzozo/go-bluetooth/bluez/profile/advertising"
)

func (app *App) GetAdvertisement() *advertising.LEAdvertisement1Properties {
	return app.advertisement
}

func (app *App) Advertise(timeout uint32) (func(), error) {

	adv := app.GetAdvertisement()

	adv.Timeout = uint16(timeout)
	adv.Duration = uint16(timeout)

	cancel, err := api.ExposeAdvertisement(app.adapterID, adv, timeout)
	return cancel, err
}
