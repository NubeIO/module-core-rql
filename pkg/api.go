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

func (inst *Module) Get(path string) ([]byte, error) {
	if path == apiRules {
		return inst.SelectAllRules()
	}

	_, uuid, combined := getPathUUID(path)
	if path == combined { // get a rule
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
	if path == apiRun {
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
