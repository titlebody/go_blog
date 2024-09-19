package desens

import "strings"

func Desensitization(tel string) string {
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "****" + tel[7:]
}

// DesensitizationEamil 邮箱脱密
func DesensitizationEamil(emali string) string {
	eList := strings.Split(emali, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****" + "@" + eList[1]
}
