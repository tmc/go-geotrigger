package geotrigger

type Device struct {
	DeviceId        string                 `json:"deviceId"`
	LastSeen        string                 `json:"lastSeen"`
	Properties      map[string]interface{} `json:"properties"`
	Tags            []string               `json:"tags"`
	TrackingProfile string                 `json:"trackingProfile"`
}

type DeviceOptions struct {
	Ids  []string
	Tags []string
	Geo  Geo
}

func (api *API) Devices(deviceOptions *DeviceOptions) ([]Device, error) {
	result := deviceResponse{}
    // @todo actually make DeviceOptions provide options
	err := api.get("/device/list", &result)
	return result.Devices, err
}

type deviceResponse struct {
	Boundingbox interface{} `json:"boundingbox"` // @todo boundingbox flesh out
	Devices     []Device    `json:"devices"`
}
