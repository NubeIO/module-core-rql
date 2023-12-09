package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-ui/backend/rumodel"
)

func (inst *Client) EdgeGetPlugins(hostUUID string) ([]rumodel.Plugin, error) {
	url := fmt.Sprintf("/host/ros/api/plugins")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&[]rumodel.Plugin{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]rumodel.Plugin)
	return *data, nil
}

func (inst *Client) EdgeGetPlugin(hostUUID, pluginName string) (*rumodel.Plugin, error) {
	url := fmt.Sprintf("/host/ros/api/plugins/name/%s", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetResult(&rumodel.Plugin{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*rumodel.Plugin), nil
}

func (inst *Client) EdgeGetConfigPlugin(hostUUID, pluginName string) (*string, error) {
	url := fmt.Sprintf("/host/ros/api/plugins/%s/config?by_plugin_name=true", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		Get(url))
	if err != nil {
		return nil, err
	}
	response := resp.String()
	return &response, nil
}

func (inst *Client) EdgeUpdateConfigPlugin(hostUUID, pluginName, config string) (*string, error) {
	url := fmt.Sprintf("/host/ros/api/plugins/%s/config?by_plugin_name=true", pluginName)
	body := map[string]string{"data": config}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	response := resp.String()
	return &response, nil
}

func (inst *Client) EdgeEnablePlugin(hostUUID, pluginName string, enable bool) (*string, error) {
	url := fmt.Sprintf("/host/ros/api/plugins/%s/enable?by_plugin_name=true", pluginName)
	body := map[string]bool{"enabled": enable}
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	state := "enabled"
	if enable == false {
		state = "disabled"
	}
	output := fmt.Sprintf("%s plugin %s successfully", pluginName, state)
	return &output, nil
}

func (inst *Client) EdgeRestartPlugin(hostUUID, pluginName string) (*string, error) {
	url := fmt.Sprintf("/host/ros/api/plugins/%s/restart?by_plugin_name=true", pluginName)
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		Post(url))
	if err != nil {
		return nil, err
	}
	output := fmt.Sprintf("%s plugin restarted successfully", pluginName)
	return &output, nil
}
