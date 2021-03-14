package src

func d2ls(d string) []string {
	switch d {
	case "2":
		return []string{"a", "b", "c"}
	case "3":
		return []string{"d", "e", "f"}
	case "4":
		return []string{"g", "h", "i"}
	case "5":
		return []string{"j", "k", "l"}
	case "6":
		return []string{"m", "n", "o"}
	case "7":
		return []string{"p", "q", "r", "s"}
	case "8":
		return []string{"t", "u", "v"}
	case "9":
		return []string{"w", "x", "y", "z"}
	default:
		return nil
	}
}

func letterCombinations(digits string) []string {
	output := make([]string, 0)

	for i, _ := range digits {
		addLst := d2ls(digits[i : i+1])
		if addLst != nil {
			outputNew := make([]string, 0)

			for _, a := range addLst {
				if len(output) == 0 {
					outputNew = append(outputNew, a)
				}

				for _, o := range output {
					outputNew = append(outputNew, o+a)
				}
			}
			output = outputNew
		}
	}
	return output
}
