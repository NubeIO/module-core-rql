package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
	"github.com/NubeIO/rubix-os/nresty"
)

func (inst *Client) FFSystemPing(hostUUID string) (*dto.Message, error) {
	url := fmt.Sprintf("/host/ros/api/system/ping")
	_, err := nresty.FormatRestyResponse(inst.Rest.R().
		SetHeader("X-Host", hostUUID).
		Get(url))
	if err != nil {
		return nil, err
	}
	return &dto.Message{Message: "ping success"}, nil
}
