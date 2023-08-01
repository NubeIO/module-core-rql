package storage

type IStorage interface {
	AddRule(rc *RQLRule) (*RQLRule, error)
	UpdateRule(uuid string, rc *RQLRule) (*RQLRule, error)
	DeleteRule(uuid string) error
	SelectRule(uuidName string) (*RQLRule, error)
	SelectAllRules() ([]RQLRule, error)
	SelectAllEnabledRules() ([]RQLRule, error)
	UpdateResult(uuid string, result interface{}) (*RQLRule, error)

	AddVariable(rc *RQLVariables) (*RQLVariables, error)
	UpdateVariable(uuid string, rc *RQLVariables) (*RQLVariables, error)
	DeleteVariable(uuid string) error
	SelectVariable(uuid string) (*RQLVariables, error)
	SelectAllVariables() ([]RQLVariables, error)
	UpdateVariableValue(uuid string, value any) (*RQLVariables, error)
	SelectVariableByName(name string) (*RQLVariables, error)
}
