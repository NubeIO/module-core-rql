package apirules

import (
	"encoding/json"
	"github.com/NubeIO/module-core-rql/helpers/float"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
)

func (inst *RQL) GetPoints(hostIDName string) any {
	resp, err := cli.GetPoints(hostIDName)
	if err != nil {
		return err
	}
	return resp
}

func (inst *RQL) GetPoint(hostIDName, uuid string) any {
	resp, err := cli.GetPoint(hostIDName, uuid)
	if err != nil {
		return err
	}
	return resp
}

func pointWriteBody(body any) (*model.Priority, error) {
	result := &model.Priority{}
	dbByte, err := json.Marshal(body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(dbByte, &result)
	return result, err
}

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

func (inst *RQL) WritePointValuePriority(hostIDName, uuid string, pri int, value float64) any {
	body, err := pointWriteBody(getPri(pri, value))
	if err != nil {
		return err
	}
	resp, err := cli.WritePointValue(hostIDName, uuid, body)
	if err != nil {
		return err
	}
	return resp
}

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
