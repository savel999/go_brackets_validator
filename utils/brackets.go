package utils

var (
	bracketsMap = map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}
)

func ValidateBrackets(str string) bool {
	var needCloseBrackets []string

	for i := 0; i < len(str); i++ {
		letter := string(str[i])
		closeBracket, isOpenBracket := bracketsMap[letter]

		if isOpenBracket {
			needCloseBrackets = append(needCloseBrackets, closeBracket)
		} else if len(needCloseBrackets) > 0 && needCloseBrackets[len(needCloseBrackets)-1] == letter {
			needCloseBrackets = needCloseBrackets[:len(needCloseBrackets)-1]
		} else {
			return false
		}

	}

	return !(len(needCloseBrackets) > 0)
}

func FixBrackets(str string) string {
	result := ""
	var needCloseBrackets []string

	for i := 0; i < len(str); i++ {
		letter := string(str[i])
		closeBracket, isOpenBracket := bracketsMap[letter]

		if isOpenBracket {
			needCloseBrackets = append(needCloseBrackets, closeBracket)
			result = result + letter
		} else if len(needCloseBrackets) > 0 {
			if needCloseBrackets[len(needCloseBrackets)-1] == letter {
				result = result + letter
			} else {
				result = result + needCloseBrackets[len(needCloseBrackets)-1]
			}

			needCloseBrackets = needCloseBrackets[:len(needCloseBrackets)-1]
		}
	}

	for i := len(needCloseBrackets); i > 0; i-- {
		result = result + needCloseBrackets[i-1]
	}

	return result
}
