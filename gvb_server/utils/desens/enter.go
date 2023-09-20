package desens

import "strings"

func DesensitizationTel(tel string) string {
	if len(tel) != 11 {
		return ""
	}
	//181 2693 4563
	return tel[:3] + "****" + tel[7:]
}

func DesensitizationEmail(email string) string {
	//122 7891082 @qq.com
	list := strings.Split(email, "@")
	if len(list) != 2 {
		return ""
	}
	return list[0][:3] + "*****@" + list[1]
}
