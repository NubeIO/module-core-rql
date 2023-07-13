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
)

const errNotFound = "not found"

func getPathUUID(path string) (urlPath, uuid, combined string) {
	s := urlSplit(path)
	if len(s) == 2 {
		return s[0], s[1], fmt.Sprintf("%s/%s", s[0], s[1])
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

func (m *Module) Get(path string) ([]byte, error) {
	if urlLen(path) == 1 { // all rules
		if path == apiRules {
			return m.SelectAllRules()
		}
	}
	if urlLen(path) == 2 {
		urlPath, uuid, combined := getPathUUID(path)
		if urlPath == apiRules {
			if path == combined { // get a rule
				return m.SelectRule(uuid)
			}
		}
	}
	return nil, errors.New(path)
}

func (m *Module) Post(path string, body []byte) ([]byte, error) {
	if path == apiRules {
		return m.AddRule(body)
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
	return nil, errors.New(errNotFound)
}

func (m *Module) Delete(path, uuid string) ([]byte, error) {
	if path == apiRules { // delete a rule
		return m.DeleteRule(uuid)
	}
	return nil, errors.New(errNotFound)
}
