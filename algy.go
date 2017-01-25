package algy

import "strings"
import "strconv"

var (
	compositMap    = make(map[rune]bool)
	numericMap     = make(map[rune]bool)
	arithMap       = make(map[rune]bool)
	operatorWeight = make(map[rune]int)
)

func buildMaps() {
	compositMap['0'] = true
	compositMap['1'] = true
	compositMap['2'] = true
	compositMap['3'] = true
	compositMap['4'] = true
	compositMap['5'] = true
	compositMap['6'] = true
	compositMap['7'] = true
	compositMap['8'] = true
	compositMap['9'] = true
	compositMap['0'] = true
	compositMap['+'] = true
	compositMap['-'] = true
	compositMap['*'] = true
	compositMap['/'] = true
	compositMap['('] = true
	compositMap[')'] = true
	compositMap['^'] = true

	numericMap['0'] = true
	numericMap['1'] = true
	numericMap['2'] = true
	numericMap['3'] = true
	numericMap['4'] = true
	numericMap['5'] = true
	numericMap['6'] = true
	numericMap['7'] = true
	numericMap['8'] = true
	numericMap['9'] = true
	numericMap['0'] = true

	arithMap['+'] = true
	arithMap['-'] = true
	arithMap['*'] = true
	arithMap['/'] = true
	arithMap['('] = true
	arithMap[')'] = true
	arithMap['^'] = true

	operatorWeight['+'] = 0
	operatorWeight['-'] = 0
	operatorWeight['*'] = 1
	operatorWeight['/'] = 1
	operatorWeight['^'] = 2
}

// Function name needs clarification for what the return means
func validateInfix(raw string) bool {
	split := strings.Split(raw, " ")

	invalid := false
	for i := range split {
		for _, e := range split[i] {
			if !compositMap[e] {
				invalid = true
			}
		}
	}

	return invalid
}

// Converts an infix string to an equivalent postfix one
func convertToPostfix(infixRaw string) string {
	infix := strings.Split(infixRaw, " ")

	var operatorStack stack
	var postfix string

	for _, r := range infix {
		for _, e := range r {
			if arithMap[e] {
				if operatorStack.Empty() {
					operatorStack.Put(string(e))
					continue
				}

				if operatorWeight[rune(operatorStack.Peek()[0])] >= operatorWeight[e] {
					postfix += operatorStack.Pop()
					postfix += " "
					operatorStack.Put(string(e))
					continue
				} else {
					operatorStack.Put(string(e))
				}
			}
		}

		if numericMap[rune(r[0])] {
			postfix += r
			postfix += " "
		}
	}

	for !operatorStack.Empty() {
		postfix += operatorStack.Pop()
		postfix += " "
	}

	// This may or may not work. Initially it trimmed the string either way,
	// this tries to add some safety to that by checking if it ends in a
	// space and accounts for it.
	if postfix[len(postfix)-1] == ' ' {
		postfix = postfix[:len(postfix)-1]
	}

	return postfix
}

func evalPostfix(raw string) string {
	var scan stack
	post := strings.Split(raw, " ")

	for _, e := range post {
		if numericMap[rune(e[0])] {
			scan.Put(string(e))
			continue
		}

		for _, r := range e {
			if arithMap[r] {
				var b int
				var a int
				var c int
				b, _ = strconv.Atoi(scan.Pop())
				a, _ = strconv.Atoi(scan.Pop())

				switch r {
				case '+':
					c = a + b
				case '-':
					c = a - b
				case '*':
					c = a * b
				case '/':
					c = a / b
				case '^':
					c = a ^ b
				}

				scan.Put(strconv.Itoa(c))
			}
		}
	}
	return scan.Peek()
}
