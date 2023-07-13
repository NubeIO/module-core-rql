package rubixoscli

import (
	"fmt"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) GetViews() ([]model.View, error) {
	url := "/api/views"

	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&[]model.View{}).
		Get(url))
	if err != nil {
		return nil, err
	}

	out := *resp.Result().(*[]model.View)
	return out, nil
}

func (inst *Client) UpdateViewName(viewIDName string, name string) (*model.View, error) {
	url := fmt.Sprintf("/api/views/%s", viewIDName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.View{}).
		SetBody(map[string]interface{}{
			"name": name,
		}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.View)
	return out, nil
}

func (inst *Client) UpdateViewTheme(viewIDName, name string, theme interface{}) (*model.View, error) {
	url := fmt.Sprintf("/api/views/%s", viewIDName)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.View{}).
		SetBody(map[string]interface{}{
			"theme": theme,
			"name":  name,
		}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.View)
	return out, nil
}

func (inst *Client) UpdateGlobalTheme(theme interface{}) (*model.View, error) {
	url := fmt.Sprintf("/api/view-settings")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.View{}).
		SetBody(map[string]interface{}{
			"theme": theme,
		}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.View)
	return out, nil
}

func (inst *Client) UpdateGlobalLogo(logo interface{}) (*model.View, error) {
	url := fmt.Sprintf("/api/view-settings")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.View{}).
		SetBody(map[string]interface{}{
			"logo": logo,
		}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.View)
	return out, nil
}

func (inst *Client) UpdateGlobalWidgetConfig(WidgetConfig interface{}) (*model.View, error) {
	url := fmt.Sprintf("/api/view-settings")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.View{}).
		SetBody(map[string]interface{}{
			"widget_config": WidgetConfig,
		}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.View)
	return out, nil
}

func (inst *Client) CreateHostView(hostUUID, name string) (*model.View, error) {
	url := fmt.Sprintf("/api/views")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.View{}).
		SetBody(map[string]interface{}{
			"host_uuid": hostUUID,
			"name":      name,
		}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.View)
	return out, nil
}

func (inst *Client) CreateHostNetworkView(groupUUID, name string) (*model.View, error) {
	url := fmt.Sprintf("/api/views")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.View{}).
		SetBody(map[string]interface{}{
			"group_uuid": groupUUID,
			"name":       name,
		}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.View)
	return out, nil
}

func (inst *Client) CreateLocationView(locationUUID, name string) (*model.View, error) {
	url := fmt.Sprintf("/api/views")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.View{}).
		SetBody(map[string]interface{}{
			"location_uuid": locationUUID,
			"name":          name,
		}).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.View)
	return out, nil
}

func (inst *Client) DeleteView(viewIDName string) (*bool, error) {
	url := fmt.Sprintf("/api/views/%s", viewIDName)

	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(false).
		Delete(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*bool)
	return out, nil
}

func (inst *Client) GetView(viewIDName string) (*model.View, error) {
	url := fmt.Sprintf("/api/views/%s", viewIDName)

	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.View{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.View)
	return out, nil
}

func (inst *Client) AddWidgetToView(viewIDName string, widget *model.ViewWidget) (*model.ViewWidget, error) {
	url := fmt.Sprintf("/api/views/widgets")

	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.ViewWidget{}).
		SetBody(widget).
		Post(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.ViewWidget)
	return out, nil
}

func (inst *Client) EditWidget(widgetUUID string, widget *model.ViewWidget) (*model.ViewWidget, error) {
	url := fmt.Sprintf("/api/views/widgets/%s", widgetUUID)
	fmt.Print("Hello, ", widget)

	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.ViewWidget{}).
		SetBody(widget).
		Patch(url))
	if err != nil {
		fmt.Print("Hello, ", err)
		return nil, err
	}
	out := resp.Result().(*model.ViewWidget)
	return out, nil
}

func (inst *Client) DeleteWidget(widgetUUID string) (*bool, error) {
	url := fmt.Sprintf("/api/views/widgets/%s", widgetUUID)

	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(false).
		Delete(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*bool)
	return out, nil
}
