package apirules

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/module-core-rql/storage"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
)

func (inst *RQL) GetVariables() any {
	out, err := inst.Storage.SelectAllVariables()
	if err != nil {
		return err
	}
	return out
}

func (inst *RQL) UpdateVariableValue(uuidName string, value any) any {
	r, err := inst.Storage.UpdateVariableValue(uuidName, value)
	if err != nil {
		return err
	}
	return r
}

/*
VarParseArray
[1, 2, "hello"]

let data = JSON.parse(RQL.VarParseArray("array"));
RQL.Return = data;
*/
func (inst *RQL) VarParseArray(uuidName string) any {
	r, err := inst.getVariable(uuidName)
	if err != nil {
		return err
	}
	b, err := json.Marshal(r.Variable)
	if err != nil {
		return err
	}
	jsonStr := string(b)
	a := gjson.Parse(jsonStr).Array()
	if len(a) > 0 {
		return a[0].String()
	}
	return a
}

/*
VarParseObject
`{"sp":22.3,"db":99.9}`

let data = RQL.VarParseObject("obj");
let sp = JSON.parse(data);
RQL.Result = sp["sp"];
*/
func (inst *RQL) VarParseObject(uuidName string) any {
	r, err := inst.getVariable(uuidName)
	if err != nil {
		return err
	}
	b, err := json.Marshal(r.Variable)
	if err != nil {
		return nil
	}
	jsonStr := string(b)
	a := gjson.Parse(jsonStr).String()
	t := strings.ReplaceAll(a, "'", "")
	t = strings.ReplaceAll(t, "`", "")
	return t
}

func (inst *RQL) VarParseString(uuidName string) any {
	r, err := inst.getVariable(uuidName)
	if err != nil {
		return err
	}
	if r == nil {
		return ""
	}
	return fmt.Sprint(r.Variable)
}

func (inst *RQL) VarParseNumber(uuidName string) any {
	r, err := inst.getVariable(uuidName)
	if err != nil {
		return err
	}
	if r == nil {
		return 0
	}
	f := r.Variable
	if s, err := strconv.ParseFloat(fmt.Sprint(f), 64); err == nil {
		return s
	}
	return 0
}

func (inst *RQL) getVariable(uuidName string) (*storage.RQLVariables, error) {
	out, err := inst.Storage.SelectAllVariables()
	if err != nil {
		return nil, err
	}
	for _, variable := range out {
		if variable.Name == uuidName {
			if err != nil {
				return &variable, err
			}
		}
		if variable.UUID == uuidName {
			if err != nil {
				return &variable, err
			}
		}
	}
	return nil, err
}

func (inst *RQL) GetVariable(uuidName string) any {
	out, err := inst.getVariable(uuidName)
	if err != nil {
		return err
	}
	return out
}
