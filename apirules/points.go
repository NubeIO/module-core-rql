package apirules

import (
	"encoding/json"
	"github.com/NubeIO/module-core-rql/helpers/float"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
)

// GetPoints retrieves points from the RubixOS API for the given host ID name.
// It makes a request to the API using the provided host ID name and returns the response.
// If an error occurs during the request, the error is returned.
// Otherwise, the response containing the points is returned.
func (inst *RQL) GetPoints(hostIDName string) any {
	resp, err := cli.GetPoints(hostIDName)
	if err != nil {
		return err
	}
	return resp
}

// GetNetworkByModuleName fetches the network by the given module name for the specified host ID.
func (inst *RQL) GetNetworkByModuleName(hostIDName, moduleName string) any {
	resp, err := cli.FFGetNetworkByPluginName(hostIDName, moduleName, true)
	if err != nil {
		return err
	}
	return resp
}

// getPointsByModuleName retrieves all points by a specific module name for a given host ID or name.
func (inst *RQL) getPointsByModuleName(hostIDName, moduleName string) ([]*model.Point, error) {
	resp, err := cli.FFGetNetworkByPluginName(hostIDName, moduleName, true)
	if err != nil {
		return nil, err
	}
	var points []*model.Point
	for _, point := range resp.Devices {
		points = append(points, point.Points...)
	}

	return points, err
}

// GetPointsByModuleName gets the points associated with a specific module name for a given host ID or name.
// It returns an error if the points cannot be retrieved.
// The function takes the hostIDName and moduleName as input parameters.
// It uses the cli.FFGetNetworkByPluginName function to fetch the data.
// The retrieved data is then processed to extract the points and return them as a slice of pointers to model.Point objects.
func (inst *RQL) GetPointsByModuleName(hostIDName, moduleName string) any {
	resp, err := cli.FFGetNetworkByPluginName(hostIDName, moduleName, true)
	if err != nil {
		return err
	}
	var points []*model.Point
	for _, point := range resp.Devices {
		points = append(points, point.Points...)
	}

	return points
}

// AllHostsPointsByModule represents the information about the points of all hosts for a specific module.
type AllHostsPointsByModule struct {
	HostUUID string
	HostName string
	Points   []*model.Point
	Error    error
	Count    int
}

// GetPointsByModuleAllHosts retrieves points by module for all hosts.
func (inst *RQL) GetPointsByModuleAllHosts(moduleName string) any {
	resp, err := cli.GetHosts()
	if err != nil {
		return err
	}
	var out []AllHostsPointsByModule
	for _, host := range resp {
		hostUUID := host.UUID
		points, err := inst.getPointsByModuleName(hostUUID, moduleName)
		newHost := AllHostsPointsByModule{
			HostUUID: host.UUID,
			HostName: host.Name,
			Points:   points,
			Error:    err,
			Count:    len(points),
		}
		out = append(out, newHost)
	}
	return out
}

// AllHostsPoints represents the points data for all hosts
type AllHostsPoints struct {
	HostUUID string
	HostName string
	Points   []model.Point
	Error    error
	Count    int
}

//
// GetPointsAllHosts calculates and retrieves the points for all hosts in the RQL instance.
//
// It makes use of the `cli` instance to retrieve the list of hosts and their respective points.
// For each host, it retrieves the points by calling the `GetPoints` method using the host's UUID.
// The points along with the host's UUID, name, error (if any), and count are stored in an instance of the `AllHostsPoints` struct.
// Finally, all the `AllHostsPoints` instances are appended to a slice and returned as the result.
//
// The function returns a slice of `AllHostsPoints` instances or an error if there was a problem retrieving the hosts.
//
func (inst *RQL) GetPointsAllHosts() any {
	resp, err := cli.GetHosts()
	if err != nil {
		return err
	}
	var out []AllHostsPoints
	for _, host := range resp {
		points, err := cli.GetPoints(host.UUID)
		newHost := AllHostsPoints{
			HostUUID: host.UUID,
			HostName: host.Name,
			Points:   points,
			Error:    err,
			Count:    len(points),
		}
		out = append(out, newHost)
	}
	return out
}

// GetPoint retrieves a point from the RQL instance using the specified host ID name and UUID.
// The point is fetched using the `cli.GetPoint` method, which returns the response and any error encountered.
// If an error occurs, it is returned as the result.
// Otherwise, the response is returned as the result.
func (inst *RQL) GetPoint(hostIDName, uuid string) any {
	resp, err := cli.GetPoint(hostIDName, uuid)
	if err != nil {
		return err
	}
	return resp
}

// pointWriteBody takes a body of type `any` and returns a pointer to model.Priority and an error.
// It serializes the body into a byte array using JSON.Marshal,
// then deserializes the byte array into the result variable using JSON.Unmarshal.
// It returns the result and any error that occurred during the process.
func pointWriteBody(body any) (*model.Priority, error) {
	result := &model.Priority{}
	dbByte, err := json.Marshal(body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(dbByte, &result)
	return result, err
}

// WritePointValue writes a point value to the specified host and UUID.
// It takes the hostIDName, uuid, and value as parameters.
// The value parameter must be of type *model.Priority.
// It returns any, which can be a *model.Priority or an error.
// If there is an error while creating the request body, it returns the error.
// If there is an error while calling cli.WritePointValue, it returns the error.
// Otherwise, it returns the response from cli.WritePointValue.
func (inst *RQL) WritePointValue(hostIDName, uuid string, value *model.Priority) any {
	body, err := pointWriteBody(value)
	if err != nil {
		return err
	}
	resp, err := cli.WritePointValue(hostIDName, uuid, body)
	if err != nil {
		return err
	}
	return resp
}

// WritePointValuePriority writes a point value with a specified priority to a host and UUID.
// It takes in the hostIDName, uuid, pri, and value parameters.
// The pri parameter determines the priority of the value.
// The value parameter represents the value to be written.
// It returns an error if there is an issue with building the point write body or writing the point value.
// If the point value is written successfully, it returns true.
//
// For the body of the point write, it uses the getPri function to create a Priority struct based on the given priority value.
// The pointWriteBody function is then used to convert the body into a Priority struct pointer.
// If there is an error during the conversion, it returns the error.
//
// Finally, it calls the WritePointValue function of the cli variable with the hostIDName, uuid, and body parameters.
// If there is an error during the write, it returns the error.
// If the write is successful, it returns true.
func (inst *RQL) WritePointValuePriority(hostIDName, uuid string, pri int, value float64) any {
	body, err := pointWriteBody(getPri(pri, value))
	if err != nil {
		return err
	}
	_, err = cli.WritePointValue(hostIDName, uuid, body)
	if err != nil {
		return err
	}
	return true
}

// getPri is a function that takes an integer `pri` and a float64 `value` as input and returns a pointer to a `model.Priority` object. Depending on the value of `pri`, different fields
func getPri(pri int, value float64) *model.Priority {
	p := &model.Priority{}
	switch pri {
	case 1:
		p = &model.Priority{
			P1: float.New(value),
		}
	case 2:
		p = &model.Priority{
			P2: float.New(value),
		}
	case 3:
		p = &model.Priority{
			P3: float.New(value),
		}
	case 4:
		p = &model.Priority{
			P4: float.New(value),
		}
	case 5:
		p = &model.Priority{
			P5: float.New(value),
		}
	case 6:
		p = &model.Priority{
			P6: float.New(value),
		}
	case 7:
		p = &model.Priority{
			P7: float.New(value),
		}
	case 8:
		p = &model.Priority{
			P8: float.New(value),
		}
	case 9:
		p = &model.Priority{
			P9: float.New(value),
		}
	case 10:
		p = &model.Priority{
			P10: float.New(value),
		}
	case 11:
		p = &model.Priority{
			P11: float.New(value),
		}
	case 12:
		p = &model.Priority{
			P12: float.New(value),
		}
	case 13:
		p = &model.Priority{
			P13: float.New(value),
		}
	case 14:
		p = &model.Priority{
			P14: float.New(value),
		}
	case 15:
		p = &model.Priority{
			P15: float.New(value),
		}
	case 16:
		p = &model.Priority{
			P16: float.New(value),
		}
	}
	return p
}
