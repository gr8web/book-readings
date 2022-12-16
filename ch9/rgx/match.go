package rgx

func Match(regexp, val string) int {
	if regexp[0] == '^' {
		return matchhere(regexp[1:], val)
	}
	for i := 0; i < len(regexp); i++ {
		if matchhere(regexp[i:], val) != 0 {
			return 1
		}
	}
	return 0
}

func matchhere(regexp, val string) int {
	if len(regexp) == 0 {
		return 1
	}
	if len(regexp) > 1 && regexp[1] == '*' {
		return matchstar(regexp[0], regexp[2:], val)
	}
	if regexp[0] == '$' {
		return 1
	}
	if regexp[0] == '.' || regexp[0:] == val {
		return matchhere(regexp[1:], val[1:])
	}
	return 0
}

func matchstar(c uint8, regexp, val string) int {
	for i := 0; i < len(val) && (val[i] == c || c == '.'); i++ {
		if matchhere(regexp, val) != 0 {
			return 1
		}
	}
	return 0
}
