package rubixoscli

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/rubix-os/nresty"
	"net/url"
)

func (inst *Client) FFGetNetworks(hostUUID string, withDevices bool, showCloneNetworks bool, overrideUrl ...string) ([]model.Network, error) {
	url := fmt.Sprintf("/host/ros/api/networks?with_tags=true&with_meta_tags=true&show_clone_networks=%t", showCloneNetworks)
	if withDevices == true {
		url = fmt.Sprintf("/host/ros/api/networks?with_devices=true&with_tags=true&with_meta_tags=true&show_clone_networks=%t", showCloneNetworks)
	}
	if buildUrl(overrideUrl...) != "" {
		url = buildUrl(overrideUrl...)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]model.Network{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	var out []model.Network
	out = *resp.Result().(*[]model.Network)
	return out, nil
}

func (inst *Client) FFGetPaginatedNetworks(hostUUID string, withDevices bool, showCloneNetworks bool, limit, offset int, search string) (*PaginationResponse, error) {
	requestURL := fmt.Sprintf("/host/ros/api/networks/paginate?with_tags=true&with_meta_tags=true&show_clone_networks=%t&limit=%v&offset=%v", showCloneNetworks, limit, offset)
	if withDevices {
		requestURL += "&with_devices=true"
	}
	if search != "" {
		requestURL += "&search_keyword=" + url.QueryEscape(search) // Ensure proper URL encoding for search value
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&PaginationResponse{}).
		Get(requestURL))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*PaginationResponse)
	return out, nil
}

func (inst *Client) FFGetNetworksWithPoints(hostUUID string) ([]model.Network, error) {
	url := fmt.Sprintf("/host/ros/api/networks?with_points=true&with_tags=true&with_meta_tags=true")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
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

func (inst *Client) FFGetNetwork(hostUUID, uuid string, withDevicesPoints bool, overrideUrl ...string) (*model.Network, error) {
	url := fmt.Sprintf("/host/ros/api/networks/%s?with_tags=true&with_meta_tags=true", uuid)
	if withDevicesPoints == true {
		url = fmt.Sprintf("/host/ros/api/networks/%s?with_devices=true&with_points=true&with_tags=true&with_meta_tags=true", uuid)
	}
	if buildUrl(overrideUrl...) != "" {
		url = buildUrl(overrideUrl...)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&model.Network{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

// FFGetNetworkWithPoints get a network with all its devices and points
func (inst *Client) FFGetNetworkWithPoints(hostUUID, uuid string) (*model.Network, error) {
	url := fmt.Sprintf("/host/ros/api/networks/%s?with_devices=true&with_points=true&with_tags=true&with_meta_tags=true", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&model.Network{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

func (inst *Client) FFDeleteNetwork(hostUUID, uuid string) (bool, error) {
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetPathParams(map[string]string{"uuid": uuid}).
		Delete("/host/ros/api/networks/{uuid}"))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (inst *Client) FFGetNetworkByPluginName(hostUUID, pluginName string, withPoints ...bool) (*model.Network, error) {
	url := fmt.Sprintf("/host/ros/api/networks/plugin-name/%s?with_tags=true&with_meta_tags=true", pluginName)
	if len(withPoints) > 0 {
		url = fmt.Sprintf("/host/ros/api/networks/plugin-name/%s?with_points=true&with_tags=true&with_meta_tags=true", pluginName)
	}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&model.Network{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

func (inst *Client) FFAddNetwork(hostUUID string, body *model.Network, restartPlugin bool) (*model.Network, error) {
	url := fmt.Sprintf("/host/ros/api/networks?restart_plugin=%t", restartPlugin)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&model.Network{}).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

func (inst *Client) FFEditNetwork(hostUUID, uuid string, body *model.Network) (*model.Network, error) {
	url := fmt.Sprintf("/host/ros/api/networks/%s", uuid)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&model.Network{}).
		SetBody(body).
		Patch(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.Network), nil
}

func (inst *Client) FFGetPollingQueueStatisticsByPluginName(hostUUID, pluginName, networkName string) (*model.PollQueueStatistics, error) {
	url := fmt.Sprintf("/host/ros/api/plugins/api/%s/polling/stats/network/name/%s", pluginName, networkName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&model.PollQueueStatistics{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*model.PollQueueStatistics), nil
}

func (inst *Client) FFGetPluginSchemaNetwork(hostUUID, pluginName string) ([]byte, error) {
	url := fmt.Sprintf("/host/ros/api/plugins/api/%s/networks/schema", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func (inst *Client) FFUpdateNetworkTag(hostUUID, networkUUID string, body []model.Tag) ([]model.Tag, error) {
	url := fmt.Sprintf("/host/ros/api/networks/%s/tags", networkUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]model.Tag{}).
		SetBody(body).
		Put(url))
	if err != nil {
		return nil, err
	}

	var out []model.Tag
	out = *resp.Result().(*[]model.Tag)
	return out, nil
}
