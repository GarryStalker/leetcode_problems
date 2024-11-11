package main

import "fmt"

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	charMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0, len(s))

	for _, elem := range s {

		_, ok := charMap[elem]
		if !ok {
			stack = append(stack, elem)
		} else {

			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] != charMap[elem] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

// TRUE				FALSE
// ()				[(])
// (){}[]			{)
// ([{}])			()[]))))
