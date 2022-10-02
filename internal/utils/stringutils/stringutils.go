package stringutils

func StartsWith(value string, c byte) bool {
	length := len(value)
	return length > 0 && value[0] == c
}

func EndsWith(value string, c byte) bool {
	length := len(value)
	return length > 0 && value[length-1] == c
}

func RemoveFirstCharacter(value string) string {
	length := len(value)
	if length == 0 {
		return value
	}

	return value[1:]
}

func RemoveLastCharacter(value string) string {
	length := len(value)
	if length == 0 {
		return value
	}

	return value[:length-1]
}
