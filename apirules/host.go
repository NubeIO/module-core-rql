package apirules

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	argspkg "github.com/NubeIO/rubix-os/args"
)

func (inst *RQL) GetHosts(args argspkg.Args) any {
	resp, err := inst.ROS.GetHosts(args)
	if err != nil {
		return err
	}
	return resp
}

func (inst *RQL) GetGroups() any {
	resp, err := cli.GetHostNetworks()
	if err != nil {
		return err
	}
	return resp
}

func (inst *RQL) GetAllHostsStatus() any {
	resp, err := cli.GetHostNetworks()
	if err != nil {
		return err
	}
	var out []*model.Group
	for _, group := range resp {
		get, err := cli.UpdateHostsStatus(group.UUID)
		if err != nil {
			return err
		}
		out = append(out, get)
	}
	return out
}

func (inst *RQL) GetHostsStatus(groupUUID string) any {
	resp, err := cli.UpdateHostsStatus(groupUUID)
	if err != nil {
		return err
	}
	resp, err = cli.GetHostNetwork(groupUUID)
	if err != nil {
		return err
	}
	return resp
}
