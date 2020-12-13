package splitstrings

func splitStrings(str string) (result []string) {
	for i := 0; i < len(str); {
		// last element
		if i != len(str)-1 {
			item := string(str[i]) + string(str[i+1])
			result = append(result, item)
		} else {
			item := string(str[i]) + string('_')
			result = append(result, item)
		}
		i = i + 2
	}
	return
}
