package codewars

import "strings"

func ToCamelCase(s string) string {
	var build strings.Builder
	flag := false
	for _, i := range s {
		if i == '-' || i == '_' {
			flag = true
			continue
		}
		if flag {
			build.WriteString(strings.ToUpper(string(i)))
			flag = false
			continue
		}
		build.WriteRune(i)
	}
	return build.String()
}
