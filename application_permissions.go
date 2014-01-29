package geotrigger

type ApplicationPermissions struct {
	DeviceList              bool `json:"deviceList"`
	DeviceLocation          bool `json:"deviceLocation"`
	DeviceTagging           bool `json:"deviceTagging"`
	DeviceToken             bool `json:"deviceToken"`
	DiscoverableApplication bool `json:"discoverableApplication"`
	DiscoverableDevice      bool `json:"discoverableDevice"`
	TriggerApply            bool `json:"triggerApply"`
	TriggerDelete           bool `json:"triggerDelete"`
	TriggerHistory          bool `json:"triggerHistory"`
	TriggerList             bool `json:"triggerList"`
	TriggerUpdate           bool `json:"triggerUpdate"`
}

func (api *API) Permissions() (*ApplicationPermissions, error) {
	result := new(ApplicationPermissions)
	err := api.get("/application/permissions", result)
	return result, err
}

func (api *API) SetPermissions(permissions *ApplicationPermissions) error {
	return api.post("/application/permissions/update", permissions)
}
