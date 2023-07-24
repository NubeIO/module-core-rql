package storage

import "github.com/NubeIO/lib-uuid/uuid"

func (inst *db) AddVariable(rc *RQLVariables) (*RQLVariables, error) {
	rc.UUID = uuid.ShortUUID("rql")
	err := inst.DB.Insert(rc)
	return rc, err
}

func (inst *db) UpdateVariable(uuid string, rc *RQLVariables) (*RQLVariables, error) {
	if rc != nil {
		rc.UUID = uuid
	}
	err := inst.DB.Update(rc)
	return rc, err
}

func (inst *db) DeleteVariable(uuid string) error {
	rule, err := inst.SelectVariable(uuid)
	if err != nil {
		return err
	}
	return inst.DB.Delete(rule)
}

func (inst *db) SelectVariable(uuid string) (*RQLVariables, error) {
	var data *RQLVariables
	err := inst.DB.Open(RQLVariables{}).Where("uuid", "=", uuid).First().AsEntity(&data)
	return data, err

}

func (inst *db) SelectAllVariables() ([]RQLVariables, error) {
	var resp []RQLVariables
	inst.DB.Open(RQLVariables{}).AsEntity(&resp)
	return resp, nil
}
