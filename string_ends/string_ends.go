package string_ends

func stringEnds(str, end string) (result bool) {
	result = false
	if len(end) <= len(str) {
		substring := str[len(str)-len(end):]
		if substring == end {
			result = true
		}
	}
	return
}
