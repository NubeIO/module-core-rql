package storage

import (
	"errors"
	"github.com/NubeIO/lib-uuid/uuid"
	"reflect"
)

func validateVariableValue(value any) error {
	if value != nil {
		kind := reflect.TypeOf(value).Kind()
		if kind == reflect.Map {
			return nil
		}
	}

	return errors.New("value must be valid json object")
}

func (inst *db) AddVariable(rc *RQLVariables) (*RQLVariables, error) {
	if err := validateVariableValue(rc.Value); err != nil {
		return nil, err
	}
	rc.UUID = uuid.ShortUUID("rql")
	err := inst.DB.Insert(rc)
	return rc, err
}

func (inst *db) UpdateVariable(uuid string, rc *RQLVariables) (*RQLVariables, error) {
	if err := validateVariableValue(rc.Value); err != nil {
		return nil, err
	}
	if rc != nil {
		rc.UUID = uuid
	}
	err := inst.DB.Update(rc)
	return rc, err
}

func (inst *db) UpdateVariableValue(uuidName string, value any) (*RQLVariables, error) {
	if err := validateVariableValue(value); err != nil {
		return nil, err
	}
	variable, err := inst.selectVariable(uuidName)
	if err != nil {
		return nil, err
	}
	if variable == nil {
		return nil, errors.New("var not found")
	}
	variable.Value = value
	return inst.UpdateVariable(variable.UUID, variable)
}

func (inst *db) DeleteVariable(uuid string) error {
	rule, err := inst.SelectVariable(uuid)
	if err != nil {
		return err
	}
	return inst.DB.Delete(rule)
}

func (inst *db) selectVariable(uuidName string) (*RQLVariables, error) {
	variable, err := inst.SelectVariable(uuidName)
	if variable == nil || err != nil {
		variable, err = inst.SelectVariableByName(uuidName)
		if err != nil {
			return nil, err
		}
	}
	return variable, err

}

func (inst *db) SelectVariable(uuid string) (*RQLVariables, error) {
	var data *RQLVariables
	err := inst.DB.Open(RQLVariables{}).Where("uuid", "=", uuid).First().AsEntity(&data)
	return data, err

}

func (inst *db) SelectVariableByName(name string) (*RQLVariables, error) {
	var data *RQLVariables
	err := inst.DB.Open(RQLVariables{}).Where("name", "=", name).First().AsEntity(&data)
	return data, err

}

func (inst *db) SelectAllVariables() ([]RQLVariables, error) {
	var resp []RQLVariables
	inst.DB.Open(RQLVariables{}).AsEntity(&resp)
	return resp, nil
}
