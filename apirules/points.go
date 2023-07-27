package apirules

import (
	"encoding/json"
	"github.com/NubeIO/module-core-rql/helpers/float"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
)

type Points struct {
	Result []model.Point `json:"result"`
	Error  string        `json:"error"`
}

type Point struct {
	Result *model.Point `json:"result"`
	Error  string       `json:"error"`
}

func (inst *RQL) GetPoints(hostIDName string) *Points {
	resp, err := cli.GetPoints(hostIDName)
	return &Points{
		Result: resp,
		Error:  errorString(err),
	}
}

func (inst *RQL) GetPoint(hostIDName, uuid string) *Point {
	resp, err := cli.GetPoint(hostIDName, uuid)
	return &Point{
		Result: resp,
		Error:  errorString(err),
	}
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

func (inst *RQL) WritePointValue(hostIDName, uuid string, value *model.Priority) *Point {
	body, err := pointWriteBody(value)
	resp, err := cli.WritePointValue(hostIDName, uuid, body)
	return &Point{
		Result: resp,
		Error:  errorString(err),
	}
}

func (inst *RQL) WritePointValuePriority(hostIDName, uuid string, pri int, value float64) *Point {
	body, err := pointWriteBody(getPri(pri, value))
	resp, err := cli.WritePointValue(hostIDName, uuid, body)
	return &Point{
		Result: resp,
		Error:  errorString(err),
	}
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
