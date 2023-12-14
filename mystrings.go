package mystrings

func Reverse(s string) string {
	result := ""
	for _, bt := range s {
		result = string(bt) + result
	}
	return result
}
