package pkg

import (
	"errors"
	"fmt"
	"strings"
)

const (
	jsonSchemaNetwork = "/schema/json/network"
	jsonSchemaDevice  = "/schema/json/device"
	jsonSchemaPoint   = "/schema/json/point"
	apiRules          = "/rules"
	apiRunExisting    = "/rules/run"
	apiRun            = "/rules/dry"
	apiVars           = "/vars"
)

const errNotFound = "not found"

func getPathUUID(path string) (urlPath, uuid, combined string) {
	s := urlSplit(path)
	if len(s) > 2 {
		return s[1], s[2], fmt.Sprintf("/%s/%s", s[1], s[2])
	}
	return "", "", ""
}

// rootPathSplit  eg /rules/run/rul_ABC123 (ROOT/SUB/NAME-UUID)
func rootPathSplit(path string) (rootPath, subPath, nameUUID string) {
	s := strings.Split(path, "/")
	s = removeEmptyStrings(s)
	if len(s) > 2 {
		return s[0], s[1], s[2]
	}
	return "", "", ""
}

func urlSplit(path string) []string {
	return strings.Split(path, "/")
}

func urlLen(path string) int {
	return len(strings.Split(path, "/"))
}

func urlGetUUID(path string) (string, error) {
	s := urlSplit(path)
	if len(s) == 0 {
		return "", errors.New("url is not correct")
	}
	lastItem := s[len(s)-1]
	if len(lastItem) < 16 {
		return "", errors.New("uuid length is incorrect")
	}
	return lastItem, nil
}

func urlIsCorrectModule(path string) bool {
	for _, s := range urlSplit(path) {
		if s == name {
			return true
		}
	}
	return false
}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func (inst *Module) Get(path string) ([]byte, error) {

	err := inst.check()
	if err != nil {
		return nil, err
	}

	if path == apiRules {
		return inst.SelectAllRules()
	}

	_, uuid, combined := getPathUUID(path)
	if path == combined { // get a rule by name or uuid http://0.0.0.0:1660/api/modules/module-core-rql/rules/test
		return inst.SelectRule(uuid)
	}

	if path == apiVars { // get all variable
		return inst.SelectAllVariables()
	}

	if path == combined { // get a variable
		return inst.SelectVariable(uuid)
	}

	return nil, errors.New(path)
}

func (inst *Module) Post(path string, body []byte) ([]byte, error) {
	if path == apiRules {
		return inst.AddRule(body)
	}

	_, subPath, nameUUID := rootPathSplit(path) // run an existing
	if subPath == "run" {
		return inst.ReuseRuleRun(body, nameUUID)
	}

	if path == apiRun { // run a rule
		return inst.Dry(body)
	}
	if path == apiVars { // add variable
		return inst.AddVariable(body)
	}

	return nil, errors.New(errNotFound)
}

func (inst *Module) Put(path, uuid string, body []byte) ([]byte, error) {
	return nil, errors.New(errNotFound)
}

func (inst *Module) Patch(path, uuid string, body []byte) ([]byte, error) {
	if path == apiRules { // update a rule
		return inst.UpdateRule(uuid, body)
	}
	if path == apiVars { // update variable
		return inst.UpdateVariable(body, uuid)
	}
	return nil, errors.New(errNotFound)
}

func (inst *Module) Delete(path, uuid string) ([]byte, error) {
	if path == apiRules { // delete a rule
		return inst.DeleteRule(uuid)
	}
	if path == apiVars { // delete a var
		return inst.DeleteVariable(uuid)
	}
	return nil, errors.New(errNotFound)
}
