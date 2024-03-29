package apirules

import (
	"github.com/NubeIO/lib-module-go/nmodule"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

func (inst *RQL) GetHosts(args nargs.Args) any {
	resp, err := inst.ROS.GetHosts(&nmodule.Opts{Args: &args})
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
