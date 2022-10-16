package util

// WrapString takes a string and limit for the number of characters before a line break,
// it returns a string that is the same as the original except for the presence of a linebreak
// after each limit characters of the input.
// Note: this implementation ignores pre-existing linebreaks
func WrapString(s string, limit int) string {
	result := ""

	for i := 0; i < len(s); i += limit {
		if len(s)-i < limit {
			result += s[i:]
			break
		}
		result += s[i:i+limit] + "\n"
	}

	return result
}
