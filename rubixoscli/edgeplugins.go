package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/nresty"
	"github.com/NubeIO/rubix-ui/backend/rumodel"
)

func (inst *Client) EdgeGetPlugins(hostIDName string) ([]rumodel.Plugin, error) {
	url := fmt.Sprintf("/proxy/ros/api/plugins")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&[]rumodel.Plugin{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	data := resp.Result().(*[]rumodel.Plugin)
	return *data, nil
}

func (inst *Client) EdgeGetPlugin(hostIDName, pluginName string) (*rumodel.Plugin, error) {
	url := fmt.Sprintf("/proxy/ros/api/plugins/path/%s", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetResult(&rumodel.Plugin{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*rumodel.Plugin), nil
}

func (inst *Client) EdgeGetConfigPlugin(hostIDName, pluginName string) (*string, error) {
	url := fmt.Sprintf("/proxy/ros/api/plugins/config/%s?by_plugin_name=true", pluginName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		Get(url))
	if err != nil {
		return nil, err
	}
	response := resp.String()
	return &response, nil
}

func (inst *Client) EdgeUpdateConfigPlugin(hostIDName, pluginName, config string) (*string, error) {
	url := fmt.Sprintf("/proxy/ros/api/plugins/config/%s?by_plugin_name=true", pluginName)
	body := map[string]string{"data": config}
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		SetBody(body).
		Post(url))
	if err != nil {
		return nil, err
	}
	response := resp.String()
	return &response, nil
}

func (inst *Client) EdgeEnablePlugin(hostIDName, pluginName string, enable bool) (*string, error) {
	url := fmt.Sprintf("/proxy/ros/api/plugins/enable/%s?by_plugin_name=true", pluginName)
	body := map[string]bool{"enabled": enable}
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
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

func (inst *Client) EdgeRestartPlugin(hostIDName, pluginName string) (*string, error) {
	url := fmt.Sprintf("/proxy/ros/api/plugins/restart/%s?by_plugin_name=true", pluginName)
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("host-uuid", hostIDName).
		SetHeader("host-name", hostIDName).
		Post(url))
	if err != nil {
		return nil, err
	}
	output := fmt.Sprintf("%s plugin restarted successfully", pluginName)
	return &output, nil
}
