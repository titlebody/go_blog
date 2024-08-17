package c_type

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1 //qq
	SignGitee SignStatus = 2 //gitee
	SignEmail SignStatus = 3 // 邮箱
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "qq"
	case SignGitee:
		str = "gitee"
	case SignEmail:
		str = "邮箱"
	default:
		str = "其他"
	}
	return str
}
