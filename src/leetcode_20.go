package src

func isMatch(c, v string) bool {
	if (c == "(" && v == ")") || (c == "[" && v == "]") || (c == "{" && v == "}") {
		return true
	}
	return false
}

func isValid(s string) bool {
	retValue := true
	stack := make([]string, 0)
	for i, _ := range s {
		v := s[i : i+1]

		switch v {
		case ")", "]", "}":
			if len(stack) == 0 {
				retValue = false
				break
			} else {
				l := len(stack)
				checkValue := stack[l-1]
				if !isMatch(checkValue, v) {
					retValue = false
					break
				} else {
					stack = stack[0 : l-1]
				}
			}
		case "(", "[", "{":
			stack = append(stack, v)
		}
	}

	if !retValue {
		return retValue
	}

	//fmt.Printf("stack:%s\n", stack)
	return len(stack) == 0
}
