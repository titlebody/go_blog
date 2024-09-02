package res

type ErrCode int

const (
	SettingsError ErrCode = 1001 //系统错误
	ArgumentError ErrCode = 1002 //参数错误
)

var (
	ErrorMap = map[ErrCode]string{
		SettingsError: "系统错误",
		ArgumentError: "参数错误",
	}
)
