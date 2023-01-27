package model

import scyna "github.com/scyna/core"

type Secret struct {
	value string
}

func ParseSecret(password string) (Secret, scyna.Error) {
	/*TODO*/
	return Secret{value: password}, nil
}

func (p Secret) Encode() string {
	return p.value /*TODO hash*/
}

func (p Secret) String() string {
	return p.value
}
