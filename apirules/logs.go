package apirules

func (inst *RQL) GetRuleLogs(uuidName string) any {
	out, err := inst.Storage.SelectRule(uuidName)
	if out == nil || err != nil {
		errMeg := "rule not found"
		if err != nil {
			errMeg = errorString(err)
		}
		return errMeg
	}
	return out.Result
}
