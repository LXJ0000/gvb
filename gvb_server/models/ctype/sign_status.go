package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1 // QQ
	SignEmail SignStatus = 2 // Email
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {

	switch s {
	case SignQQ:
		return "QQ"
	case SignEmail:
		return "Email"
	default:
		return "其他"
	}
}
