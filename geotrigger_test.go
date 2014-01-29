package geotrigger_test

import (
	"fmt"
	"log"

	geotrigger "github.com/tmc/go-geotrigger"
)

var (
	clientID, clientSecret = "z0hJO9POX7FWd2IA", "aaec844e0541495ca6ee7b0a1f303e5f"
)

func ExampleNewAPI_changingPermissions() {
	api, err := geotrigger.NewAPI(clientID, clientSecret)
	if err != nil {
		log.Println("error connecting to geotrigger api:", err)
		return
	}
	p, err := api.Permissions()
	if err != nil {
		panic(err)
	}
	fmt.Println("TriggerUpdate:", p.TriggerUpdate)

	p.TriggerUpdate = !p.TriggerUpdate
	if err = api.SetPermissions(p); err != nil {
		panic(err)
	}

	p, err = api.Permissions()
	if err != nil {
		panic(err)
	}
	fmt.Println("TriggerUpdate:", p.TriggerUpdate)

	// Output:
	// foo
}

func ExampleNewAPI_listDevices() {
	api, err := geotrigger.NewAPI(clientID, clientSecret)
	if err != nil {
		log.Println("error connecting to geotrigger api:", err)
		return
	}
	devices, err := api.Devices(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("devices: %+v\n", devices)
	// Output:
	// foo
}
