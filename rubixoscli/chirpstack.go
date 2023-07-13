package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-ui/backend/chirpstack"
)

const limit = "2000"

// most nube supported sensors are now added as OTA devices

// CSLogin to CS with username and password to get token if not provided in config
func (inst *Client) CSLogin(hostIDName, user, pass string) (string, error) {
	token := &CSLoginToken{}
	url := "/proxy/chirp/api/internal/login"
	_, err := inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
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
func (inst *Client) CSGetAdminTokens(hostIDName, token string) (*chirpstack.GetTokens, error) {
	q := fmt.Sprintf("/proxy/chirp/api/internal/api-keys?limit=30&isAdmin=true")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.GetTokens{}).
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetHeader("cs-token", token).
		Get(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.GetTokens), nil
}

func (inst *Client) CSAddAdminToken(hostIDName, token string, body *chirpstack.AdminToken) (*chirpstack.AdminTokenResponse, error) {
	q := fmt.Sprintf("/proxy/chirp/api/internal/api-keys")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.AdminTokenResponse{}).
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
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
func (inst *Client) CSGetApplications(hostIDName string) (*chirpstack.Applications, error) {
	q := fmt.Sprintf("/proxy/ros/api/plugins/api/lorawan/cs/applications=%s", limit)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.Applications{}).
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		Get(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.Applications), nil
}

// CSGetGateways get all gateways
func (inst *Client) CSGetGateways(hostIDName string) (*chirpstack.Gateways, error) {
	q := fmt.Sprintf("/proxy/ros/api/plugins/api/lorawan/cs/gateways?limit=%s", limit)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.Gateways{}).
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		Get(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.Gateways), nil
}

// CSGetDevices get all
func (inst *Client) CSGetDevices(hostIDName, applicationID string) (*chirpstack.Devices, error) {
	q := fmt.Sprintf("/proxy/ros/api/plugins/api/lorawan/cs/devices?limit=%s&applicationID=%s", limit, applicationID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.Devices{}).
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		Get(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.Devices), nil
}

// CSGetDevice get a device
func (inst *Client) CSGetDevice(hostIDName, devEui string) (*chirpstack.Device, error) {
	q := fmt.Sprintf("/proxy/ros/api/plugins/api/lorawan/cs/devices/%s", devEui)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.Device{}).
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		Get(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.Device), nil
}

// CSGetDeviceProfiles get all
func (inst *Client) CSGetDeviceProfiles(hostIDName string) (*chirpstack.DeviceProfiles, error) {
	q := fmt.Sprintf("/proxy/ros/api/plugins/api/lorawan/cs/device-profiles?limit=%s", limit)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.DeviceProfiles{}).
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		Get(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.DeviceProfiles), nil
}

// CSAddDevice add all
func (inst *Client) CSAddDevice(hostIDName string, body *chirpstack.Device) (*chirpstack.Device, error) {
	q := fmt.Sprintf("/proxy/ros/api/plugins/api/lorawan/cs/devices")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.Device{}).
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		Post(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.Device), nil
}

// CSEditDevice edit object
func (inst *Client) CSEditDevice(hostIDName, devEui string, body *chirpstack.Device) (*chirpstack.Device, error) {
	q := fmt.Sprintf("/proxy/ros/api/plugins/api/lorawan/cs/devices/%s/", devEui)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.Device{}).
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		Put(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.Device), nil
}

// CSDeleteDevice delete
func (inst *Client) CSDeleteDevice(hostIDName, devEui string) (bool, error) {
	q := fmt.Sprintf("/proxy/ros/api/plugins/api/lorawan/cs/devices/%s", devEui)
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		Delete(q))
	if err != nil {
		return false, err
	}
	return true, nil
}

// CSDeviceOTAKeysUpdate active a device
func (inst *Client) CSDeviceOTAKeysUpdate(hostIDName, devEui string, body *chirpstack.DeviceKey) (*chirpstack.DeviceKey, error) {
	q := fmt.Sprintf("/proxy/ros/api/plugins/api/lorawan/cs/devices/keys/%s", devEui)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.DeviceKey{}).
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		Put(q))
	if err != nil {
		return nil, err
	}
	r := resp.Result().(*chirpstack.DeviceKey)
	return r, nil
}

// CSDeviceOTAKeys active a device
func (inst *Client) CSDeviceOTAKeys(hostIDName, devEui string, body *chirpstack.DeviceKey) (*chirpstack.DeviceKey, error) {
	q := fmt.Sprintf("/proxy/ros/api/plugins/api/lorawan/cs/devices/keys/%s", devEui)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.DeviceKey{}).
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		Post(q))
	if err != nil {
		return nil, err
	}
	r := resp.Result().(*chirpstack.DeviceKey)
	return r, nil
}

// CSActivateDevice active a device
func (inst *Client) CSActivateDevice(hostIDName, devEui string, body *chirpstack.DeviceActivation) (*chirpstack.DeviceActivation, error) {
	q := fmt.Sprintf("/proxy/ros/api/plugins/api/lorawan/cs/devices/activate/%s", devEui)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(chirpstack.DeviceActivation{}).
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		Put(q))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*chirpstack.DeviceActivation), nil
}
