package apirules

func (inst *RQL) GetModule(hostIDName, moduleName string) any {
	resp, err := cli.EdgeGetPlugin(hostIDName, moduleName)
	if err != nil {
		return err
	}
	return resp
}
