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

func (m *Module) Get(path string) ([]byte, error) {

	err := m.check()
	if err != nil {
		return nil, err
	}

	if path == apiRules {
		return m.SelectAllRules()
	}

	_, uuid, combined := getPathUUID(path)
	if path == combined { // get a rule by name or uuid http://0.0.0.0:1660/api/modules/module-core-rql/rules/test
		return m.SelectRule(uuid)
	}

	if path == apiVars { // get all variable
		return m.SelectAllVariables()
	}

	if path == combined { // get a variable
		return m.SelectVariable(uuid)
	}

	return nil, errors.New(path)
}

func (m *Module) Post(path string, body []byte) ([]byte, error) {
	if path == apiRules {
		return m.AddRule(body)
	}

	_, subPath, nameUUID := rootPathSplit(path) // run an existing
	if subPath == "run" {
		return m.ReuseRuleRun(body, nameUUID)
	}

	if path == apiRun { // run a rule
		return m.Dry(body)
	}
	if path == apiVars { // add variable
		return m.AddVariable(body)
	}

	return nil, errors.New(errNotFound)
}

func (m *Module) Put(path, uuid string, body []byte) ([]byte, error) {
	return nil, errors.New(errNotFound)
}

func (m *Module) Patch(path, uuid string, body []byte) ([]byte, error) {
	if path == apiRules { // update a rule
		return m.UpdateRule(uuid, body)
	}
	if path == apiVars { // update variable
		return m.UpdateVariable(body, uuid)
	}
	return nil, errors.New(errNotFound)
}

func (m *Module) Delete(path, uuid string) ([]byte, error) {
	if path == apiRules { // delete a rule
		return m.DeleteRule(uuid)
	}
	if path == apiVars { // delete a var
		return m.DeleteVariable(uuid)
	}
	return nil, errors.New(errNotFound)
}
