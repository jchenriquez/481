package integertoenglish

import (
	"fmt"
	"math"
	"strings"
)

var ones = map[int]string{
	0: "",
	1: "One",
	2: "Two",
	3: "Three",
	4: "Four",
	5: "Five",
	6: "Six",
	7: "Seven",
	8: "Eight",
	9: "Nine",
}

var onesStrToNum = map[string]int{
	"One":   1,
	"Two":   2,
	"Three": 3,
	"Four":  4,
	"Five":  5,
	"Six":   6,
	"Seven": 7,
	"Eight": 8,
	"Nine":  9,
}

var teens = map[int]string{
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
}

var tens = map[int]string{
	2: "Twenty",
	3: "Thirty",
	4: "Forty",
	5: "Fifty",
	6: "Sixty",
	7: "Seventy",
	8: "Eighty",
	9: "Ninety",
}

var tenPowersPositions = map[int]string{
	0: "",
	1: "Thousand",
	2: "Million",
	3: "Billion",
	4: "Trillion",
}

func hundred(num int) (ret string) {

	var digitIndex int

	for num > 0 {
		floated := float64(num) / 10
		notFloated := num / 10
		digit := int(math.Round((floated - float64(notFloated)) * 10))
		ret = strings.Trim(ret, " ")

		switch digitIndex {
		case 0:
			one, _ := ones[digit]
			ret = fmt.Sprintf(" %s %s", one, ret)
		case 1:
			if digit == 1 && len(ret) == 0 {
				ret = teens[10]
			} else if digit == 1 {
				inDigit := onesStrToNum[ret]
				teen := teens[10+inDigit]

				ret = teen
			} else {
				ten := tens[digit]

				ret = fmt.Sprintf(" %s %s", ten, ret)
			}
		case 2:
			one := ones[digit]

			ret = fmt.Sprintf(" %s Hundred %s", one, ret)
		}

		digitIndex++
		num = notFloated
	}

	ret = strings.Trim(ret, " ")

	return
}

// NumberToWords will take num and return the worded representation of that number
func NumberToWords(num int) (ret string) {
	if num == 0 {
		return "Zero"
	}

	mover := math.Pow(10, 3)
	tenPowerPos := 0

	for num > 0 {
		floated := float64(num) / mover
		notFloated := num / int(mover)
		digits := int(math.Round((floated - float64(notFloated)) * mover))
		ret = strings.Trim(ret, " ")

		if digits > 0 {
			hundr := hundred(digits)
			tensPwer := tenPowersPositions[tenPowerPos]
			ret = fmt.Sprintf(" %s %s %s", hundr, tensPwer, ret)
		}

		tenPowerPos++
		num = notFloated
	}

	ret = strings.Trim(ret, " ")

	return
}
