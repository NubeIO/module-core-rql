package apirules

import (
	"github.com/NubeIO/module-core-rql/storage"
)

type RuleLogsResponse struct {
	Result []storage.Result
	Error  string
}

func (inst *RQL) GetRuleLogs(uuidName string) *RuleLogsResponse {
	out, err := inst.Storage.SelectRule(uuidName)
	if out == nil || err != nil {
		errMeg := "rule not found"
		if err != nil {
			errMeg = errorString(err)
		}
		return &RuleLogsResponse{
			Result: nil,
			Error:  errMeg,
		}
	}
	return &RuleLogsResponse{
		Result: out.Result,
		Error:  errorString(err),
	}
}
