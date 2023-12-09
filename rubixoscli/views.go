package rubixoscli

import (
	"fmt"

	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) GetViews() ([]model.View, error) {
	url := "/api/views?with_widgets=true"

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

func (inst *Client) GetViewSettings() (*model.ViewSetting, error) {
	url := fmt.Sprintf("/api/view-settings")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.ViewSetting{}).
		Get(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.ViewSetting)
	return out, nil
}

func (inst *Client) UpdateGlobalTheme(theme interface{}) (*model.View, error) {
	url := fmt.Sprintf("/api/view-settings")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.View{}).
		SetBody(UpdateGlobalLogo{
			Theme: theme,
		}).
		Put(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.View)
	return out, nil
}

type UpdateGlobalLogo struct {
	Theme interface{} `json:"theme"`
}

func (inst *Client) UpdateGlobalLogo(logo interface{}) (*model.ViewSetting, error) {
	url := fmt.Sprintf("/api/view-settings")
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.ViewSetting{}).
		SetBody(map[string]interface{}{
			"logo": logo,
		}).
		Put(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.ViewSetting)
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
	url := fmt.Sprintf("/api/views/%s?with_widgets=true&widget_order=ASC", viewIDName)

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

	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(&model.ViewWidget{}).
		SetBody(widget).
		Patch(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*model.ViewWidget)
	return out, nil
}

type OrderRequestBody struct {
	Order int `json:"order"`
}

func (inst *Client) UpdateWidgetOrder(widgetUUID string, order int) (*bool, error) {
	url := fmt.Sprintf("/api/views/widgets/%s/order", widgetUUID)
	resp, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetResult(false).
		SetBody(OrderRequestBody{
			Order: order,
		}).
		Patch(url))
	if err != nil {
		return nil, err
	}
	out := resp.Result().(*bool)
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
