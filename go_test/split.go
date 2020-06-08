package split

import "strings"

func Split(str string, sep string) []string {
	index := strings.Index(str, sep)
	//res := make([]string, 0)
	res := make([]string, 0, strings.Count(str, sep) + 1)//优化
	for index >= 0 {
		res = append(res, string(str[:index]))
		str = str[index + len(sep):]
		index = strings.Index(str, sep)
	}
	res = append(res, str)
	return res
}
