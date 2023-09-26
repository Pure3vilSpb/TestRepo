package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	pattern            = `^(I|(II)|(III)|(IV)|V|(VI)|(VII)|(VIII)|(IX)|X)[\+\*\/\-](I|(II)|(III)|(IV)|V|(VI)|(VII)|(VIII)|(IX)|X)$|^(([1-9]|10)[\+\*\/\-]([1-9]|10))$`
	patternObj         = regexp.MustCompile(pattern)
	reader             = bufio.NewReader(os.Stdin)
	romanArabMap       = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	input              string
	op                 byte
	convNum1, convNum2 int
	result             int
)

func convertRomanToArab(romNum string) (int, error) {
	for key, value := range romanArabMap {
		if romNum == key {
			return value, nil
		}
	}
	return 0, errors.New(`the number is outside of the allowed range`)
}

// func convertArabToRoman(arabNum int) string {
// 	for key, value := range romanArabMap {
// 		if arabNum == value {
// 			return key
// 		}
// 	}
// 	return ""
// }

func convertArabToRoman(arabNum int) (string, error) {
	if arabNum < 1 {
		return "", errors.New(`error! There is no zero or negative numbers in Roman numeral system`)
	}

	romanNumerals := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	arabicValues := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	var convResult string
	for i := 0; i < len(arabicValues); i++ {
		for arabNum >= arabicValues[i] {
			arabNum -= arabicValues[i]
			convResult += romanNumerals[i]
		}
	}

	return convResult, nil
}

func calculate(operator byte, numb1, numb2 int) (int, error) {
	switch operator {
	case '+':
		return numb1 + numb2, nil
	case '-':
		return numb1 - numb2, nil
	case '*':
		return numb1 * numb2, nil
	case '/':
		return numb1 / numb2, nil
	default:
		return 0, errors.New(`invalid operator`)
	}
}

func main() {

	for {
		fmt.Println("Please enter a math expression or \"exit\" to quit ")

		input, _ = reader.ReadString('\n')
		//input = "V*X"

		input = strings.ReplaceAll(input, " ", "")
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting program...")
			break
		} else if !(patternObj.MatchString(input)) {
			fmt.Println("Only +-*/ expressions of roman and arabic numbers from 1 to 10 are allowed")
			continue
		}

		opIndex := strings.IndexAny(input, "+-*/")

		op = input[opIndex]

		num1 := input[0:opIndex]
		num2 := input[opIndex+1:]

		if !(regexp.MustCompile(`^\d+$`).MatchString(num1)) {
			convNum1, err1 := convertRomanToArab(num1)
			convNum2, err2 := convertRomanToArab(num2)
			if err1 != nil {
				fmt.Println(err1.Error())
				break
			} else if err2 != nil {
				fmt.Println(err2.Error())
				break
			} else {
				result, calcErr := calculate(op, convNum1, convNum2)

				if calcErr != nil {
					fmt.Println(calcErr.Error())
					break
				} else {
					convResult, err := convertArabToRoman(result)
					if err != nil {
						fmt.Printf("The expression: %v %c %v is invalid\n%s\n ", num1, op, num2, err.Error())
						break
					} else {
						fmt.Printf("%v %T %c %v %T = %v\n", num1, num1, op, num2, num2, convResult)
					}
				}

			}
		} else {
			convNum1, err1 := strconv.Atoi(input[0:opIndex])
			convNum2, err2 := strconv.Atoi(input[opIndex+1:])
			if err1 != nil {
				fmt.Println(err1.Error())
				break
			} else if err2 != nil {
				fmt.Println(err2.Error())
				break
			} else {
				result, calcErr := calculate(op, convNum1, convNum2)

				if calcErr != nil {
					fmt.Println(calcErr.Error())
					break
				} else {
					fmt.Printf("%v %c %v = %v\n", num1, op, num2, result)
				}
			}
		}

	}
}
