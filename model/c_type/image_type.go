package c_type

import "encoding/json"

type ImageType int

const (
	Local ImageType = 1 //本地
	QiNiu ImageType = 2 //七牛
)

func (s ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageType) String() string {
	var str string
	switch s {
	case Local:
		str = "本地"
	case QiNiu:
		str = "七牛"
	default:
		str = "其他"
	}
	return str
}
