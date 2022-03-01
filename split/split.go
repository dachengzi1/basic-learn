package split

import "strings"

func Split(str, s string) (result []string) {
	i := strings.Index(str, s)
	for i > -1 {
		result = append(result, str[:i])
		str = str[len(s)+1:]
		i = strings.Index(str, s)
	}
	result = append(result, str)
	return
}
