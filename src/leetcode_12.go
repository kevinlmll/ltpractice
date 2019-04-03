package src

import "fmt"

func intToRoman(num int) string {
	thousands := []string{"M", "MM", "MMM"}
	hundreds := []string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	output := ""
	if num/1000 > 0 {
		a := num / 1000
		output = output + thousands[a-1]
		num = num % 1000
	}
	if num/100 > 0 {
		a := num / 100
		output = output + hundreds[a-1]
		num = num % 100
	}
	if num/10 > 0 {
		a := num / 10
		output = output + tens[a-1]
		num = num % 10
	}
	if num > 0 {
		output = output + ones[num-1]
	}
	return output
}

func main() {
	fmt.Printf("intToRoman:%s\n", intToRoman(3999))
}
