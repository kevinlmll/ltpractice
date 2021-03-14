package src

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}

	//1. find first not space charactor
	i := 0
	l := len(str)
	filter := 0

	for i < l {
		if str[i] != ' ' {
			break
		}
		i = i + 1
	}
	if i == l {
		return 0
	}

	switch str[i] {
	case '-', '+':
		filter = 1
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		filter = 1
	default:
		filter = 0
	}

	if filter == 0 {
		return 0
	}

	o := int64(0)
	sign := 1
	if str[i] == '-' {
		sign = -1
		i = i + 1
	} else if str[i] == '+' {
		sign = 1
		i = i + 1
	}

	for i < l {
		if str[i] != '0' {
			break
		}
		i = i + 1
	}

	invalid := 0
	digitNum := 0
	min := int64(-(1 << 31))
	max := int64((1 << 31) - 1)
	for i < l {
		switch str[i] {
		case '0':
			o = o*10 + 0
		case '1':
			o = o*10 + 1
		case '2':
			o = o*10 + 2
		case '3':
			o = o*10 + 3
		case '4':
			o = o*10 + 4
		case '5':
			o = o*10 + 5
		case '6':
			o = o*10 + 6
		case '7':
			o = o*10 + 7
		case '8':
			o = o*10 + 8
		case '9':
			o = o*10 + 9
		default:
			invalid = 1
		}

		if invalid == 1 {
			break
		}
		i = i + 1
		digitNum = digitNum + 1

		if digitNum > 10 {
			break
		}
	}

	o = o * int64(sign)
	if o < min {
		o = min
	} else if o > max {
		o = max
	}

	return int(o)
}
