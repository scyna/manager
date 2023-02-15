package model

type Module struct {
	Code     string
	secret   string
	Sessions []Session
}

func NewModule(code string, secret string) (Module, error) {
	// TODO: check valid code name, secret value
	return Module{Code: code, secret: secret}, nil
}

func (m Module) MatchSecret(secret string) bool {
	return m.secret == secret
}
