package rubixoscli

import (
	"errors"
	"fmt"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) FFGetNetworks(hostIDName string, withDevices bool, showCloneNetworks bool, overrideUrl ...string) ([]model.Network, error) {
	url := fmt.Sprintf("/proxy/ros/api/networks?with_tags=true&with_meta_tags=true&show_clone_networks=%t", showCloneNetworks)
	if withDevices == true {
		url = fmt.Sprintf("/proxy/ros/api/networks?with_devices=true&with_tags=true&with_meta_tags=true&show_clone_networks=%t", showCloneNetworks)
	}
	if buildUrl(overrideUrl...) != "" {
		url = buildUrl(overrideUrl...)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]model.Network{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Network
	out = *resp.Result().(*[]model.Network)
	return out, nil
}

func (inst *Client) FFGetNetworksWithPoints(hostIDName string) ([]model.Network, error) {
	url := fmt.Sprintf("/proxy/ros/api/networks?with_points=true&with_tags=true&with_meta_tags=true")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]model.Network{}).
		Get(url))
	if err != nil {
		if err.Error() == "<nil>" {
			return nil, errors.New(resp.Status())
		}
		return nil, err
	}
	if resp.StatusCode() > 300 {
		return nil, errors.New(resp.Status())
	}
	var out []model.Network
	out = *resp.Result().(*[]model.Network)
	return out, nil
}

func (inst *Client) FFGetNetwork(hostIDName, uuid string, withDevicesPoints bool, overrideUrl ...string) (*model.Network, error) {
	url := fmt.Sprintf("/proxy/ros/api/networks/%s?with_tags=true&with_meta_tags=true", uuid)
	if withDevicesPoints == true {
		url = fmt.Sprintf("/proxy/ros/api/networks/%s?with_devices=true&with_points=true&with_tags=true&with_meta_tags=true", uuid)
	}
	if buildUrl(overrideUrl...) != "" {
		url = buildUrl(overrideUrl...)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Network{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

// FFGetNetworkWithPoints get a network with all its devices and points
func (inst *Client) FFGetNetworkWithPoints(hostIDName, uuid string) (*model.Network, error) {
	url := fmt.Sprintf("/proxy/ros/api/networks/%s?with_devices=true&with_points=true&with_tags=true&with_meta_tags=true", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Network{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

func (inst *Client) FFDeleteNetwork(hostIDName, uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("/proxy/ros/api/networks/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (inst *Client) GetNetworkByPluginName(hostIDName, pluginName string, withPoints ...bool) (*model.Network, error) {
	url := fmt.Sprintf("/proxy/ros/api/networks/plugin/%s?with_tags=true&with_meta_tags=true", pluginName)
	if len(withPoints) > 0 {
		url = fmt.Sprintf("/proxy/ros/api/networks/plugin/%s?with_points=true&with_tags=true&with_meta_tags=true", pluginName)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Network{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

func (inst *Client) FFAddNetwork(hostIDName string, body *model.Network, restartPlugin bool) (*model.Network, error) {
	url := fmt.Sprintf("/proxy/ros/api/networks?restart_plugin=%t", restartPlugin)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Network{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

func (inst *Client) FFEditNetwork(hostIDName, uuid string, body *model.Network) (*model.Network, error) {
	url := fmt.Sprintf("/proxy/ros/api/networks/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.Network{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

func (inst *Client) FFGetPollingQueueStatisticsByPluginName(hostIDName, pluginName, networkName string) (*model.PollQueueStatistics, error) {
	url := fmt.Sprintf("/proxy/ros/api/plugins/api/%s/polling/stats/network/%s", pluginName, networkName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&model.PollQueueStatistics{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.PollQueueStatistics), nil
}

func (inst *Client) FFGetPluginSchemaNetwork(hostIDName, pluginName string) ([]byte, error) {
	url := fmt.Sprintf("/proxy/ros/api/plugins/api/%s/schema/json/network", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
