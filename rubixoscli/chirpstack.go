package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-ui/backend/chirpstack"
	"strings"
)

const limit = "2000"

// most nube supported sensors are now added as OTA devices

// CSLogin to CS with username and password to get token if not provided in config
func (inst *Client) CSLogin(hostUUID, user, pass string) (string, error) {
	token := &CSLoginToken{}
	url := "/chirp/api/internal/login"
	_, err := inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(CSCredentials{
			Email:    user,
			Password: pass,
		}).
		SetResult(&token).
		Post(url)
	if token != nil {
		return token.Token, err
	}
	return "", err
}

// CSGetAdminTokens get all https://github.com/NubeIO/rubix-ce/issues/890
func (inst *Client) CSGetAdminTokens(hostUUID, token string) (*chirpstack.GetTokens, error) {
	q := fmt.Sprintf("/chirp/api/internal/api-keys?limit=30&isAdmin=true")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.GetTokens{}).
		SetHeader("X-Host", hostUUID).
		SetHeader("cs-token", token).
		Get(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.GetTokens), nil
}

func (inst *Client) CSAddAdminToken(hostUUID, token string, body *chirpstack.AdminToken) (*chirpstack.AdminTokenResponse, error) {
	q := fmt.Sprintf("/chirp/api/internal/api-keys")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.AdminTokenResponse{}).
		SetHeader("X-Host", hostUUID).
		SetHeader("cs-token", token).
		SetBody(body).
		Post(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.AdminTokenResponse), nil
}

type CSApplications struct {
	Result []struct {
		ID string `json:"id"`
	} `json:"result"`
}

type CSCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CSLoginToken struct {
	Token string `json:"jwt"`
}

// CSGetApplications get all
func (inst *Client) CSGetApplications(hostUUID, pluginName string) (*chirpstack.Applications, error) {
	var q string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		q = fmt.Sprintf("/host/ros/api/modules/module-core-lorawan/cs/applications=%s", limit)
	} else {
		q = fmt.Sprintf("/host/ros/api/plugins/api/lorawan/cs/applications=%s", limit)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.Applications{}).
		SetHeader("X-Host", hostUUID).
		Get(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.Applications), nil
}

// CSGetGateways get all gateways
func (inst *Client) CSGetGateways(hostUUID, pluginName string) (*chirpstack.Gateways, error) {
	var q string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		q = fmt.Sprintf("/host/ros/api/modules/module-core-lorawan/cs/gateways?limit=%s", limit)
	} else {
		q = fmt.Sprintf("/host/ros/api/plugins/api/lorawan/cs/gateways?limit=%s", limit)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.Gateways{}).
		SetHeader("X-Host", hostUUID).
		Get(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.Gateways), nil
}

// CSGetDevices get all
func (inst *Client) CSGetDevices(hostUUID, applicationID, pluginName string) (*chirpstack.Devices, error) {
	var q string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		q = fmt.Sprintf("/host/ros/api/modules/module-core-lorawan/cs/devices?limit=%s&applicationID=%s", limit, applicationID)
	} else {
		q = fmt.Sprintf("/host/ros/api/plugins/api/lorawan/cs/devices?limit=%s&applicationID=%s", limit, applicationID)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.Devices{}).
		SetHeader("X-Host", hostUUID).
		Get(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.Devices), nil
}

// CSGetDevice get a device
func (inst *Client) CSGetDevice(hostUUID, devEui, pluginName string) (*chirpstack.Device, error) {
	var q string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		q = fmt.Sprintf("/host/ros/api/modules/module-core-lorawan/cs/devices/%s", devEui)
	} else {
		q = fmt.Sprintf("/host/ros/api/plugins/api/lorawan/cs/devices/%s", devEui)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.Device{}).
		SetHeader("X-Host", hostUUID).
		Get(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.Device), nil
}

// CSGetDeviceProfiles get all
func (inst *Client) CSGetDeviceProfiles(hostUUID, pluginName string) (*chirpstack.DeviceProfiles, error) {
	var q string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		q = fmt.Sprintf("/host/ros/api/modules/module-core-lorawan/cs/device-profiles?limit=%s", limit)
	} else {
		q = fmt.Sprintf("/host/ros/api/plugins/api/lorawan/cs/device-profiles?limit=%s", limit)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.DeviceProfiles{}).
		SetHeader("X-Host", hostUUID).
		Get(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.DeviceProfiles), nil
}

// CSAddDevice add all
func (inst *Client) CSAddDevice(hostUUID, pluginName string, body *chirpstack.Device) (*chirpstack.Device, error) {
	var q string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		q = fmt.Sprintf("/host/ros/api/modules/module-core-lorawan/cs/devices")
	} else {
		q = fmt.Sprintf("/host/ros/api/plugins/api/lorawan/cs/devices")
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.Device{}).
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		Post(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.Device), nil
}

// CSEditDevice edit object
func (inst *Client) CSEditDevice(hostUUID, devEui, pluginName string, body *chirpstack.Device) (*chirpstack.Device, error) {
	var q string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		q = fmt.Sprintf("/host/ros/api/modules/module-core-lorawan/cs/devices/%s", devEui)
	} else {
		q = fmt.Sprintf("/host/ros/api/plugins/api/lorawan/cs/devices/%s", devEui)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.Device{}).
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		Put(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.Device), nil
}

// CSDeleteDevice delete
func (inst *Client) CSDeleteDevice(hostUUID, devEui, pluginName string) (bool, error) {
	var q string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		q = fmt.Sprintf("/host/ros/api/modules/module-core-lorawan/cs/devices/%s", devEui)
	} else {
		q = fmt.Sprintf("/host/ros/api/plugins/api/lorawan/cs/devices/%s", devEui)
	}
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		Delete(q))
	if err != nil {
		return false, err
	}
	return true, nil
}

// CSDeviceOTAKeysUpdate active a device
func (inst *Client) CSDeviceOTAKeysUpdate(hostUUID, devEui, pluginName string, body *chirpstack.DeviceKey) (*chirpstack.DeviceKey, error) {
	var q string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		q = fmt.Sprintf("/host/ros/api/modules/module-core-lorawan/cs/devices/keys/%s", devEui)
	} else {
		q = fmt.Sprintf("/host/ros/api/plugins/api/lorawan/cs/devices/keys/%s", devEui)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.DeviceKey{}).
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		Put(q))
	if err != nil {
		return nil, err
	}
	r := resp.Result().(*chirpstack.DeviceKey)
	return r, nil
}

// CSDeviceOTAKeys active a device
func (inst *Client) CSDeviceOTAKeys(hostUUID, devEui, pluginName string, body *chirpstack.DeviceKey) (*chirpstack.DeviceKey, error) {
	var q string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		q = fmt.Sprintf("/host/ros/api/modules/module-core-lorawan/cs/devices/keys/%s", devEui)
	} else {
		q = fmt.Sprintf("/host/ros/api/plugins/api/lorawan/cs/devices/keys/%s", devEui)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.DeviceKey{}).
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		Post(q))
	if err != nil {
		return nil, err
	}
	r := resp.Result().(*chirpstack.DeviceKey)
	return r, nil
}

// CSActivateDevice active a device
func (inst *Client) CSActivateDevice(hostUUID, devEui, pluginName string, body *chirpstack.DeviceActivation) (*chirpstack.DeviceActivation, error) {
	var q string
	if strings.HasPrefix(pluginName, ModulePrefix) {
		q = fmt.Sprintf("/host/ros/api/modules/module-core-lorawan/cs/devices/activate/%s", devEui)
	} else {
		q = fmt.Sprintf("/host/ros/api/plugins/api/lorawan/cs/devices/activate/%s", devEui)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.DeviceActivation{}).
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		Put(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.DeviceActivation), nil
}
