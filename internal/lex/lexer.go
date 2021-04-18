package lex

import (
	"regexp"
	"strings"
)

var Chars []string = []string{}
var Pointer int = -1
var lastChar string = ""

func GetChar() {
	if Pointer+1 > len(Chars)-1 {
		lastChar = "EOF"
	} else {
		Pointer += 1
		lastChar = Chars[Pointer]
	}
}
func Peek() string {
	if Pointer == len(Chars)-1 {
		return "EOF"
	} else {
		return Chars[Pointer+1]
	}
}
func isSpace() bool {
	space, _ := regexp.MatchString(`\s`, lastChar)
	return space
}
func isAlpha() bool {
	alpha, _ := regexp.MatchString(`[A-Za-z_]`, lastChar)
	return alpha
}
func isAlphaNum(exp ...string) bool {
	if len(exp) != 0 {
		alpha, _ := regexp.MatchString(`[A-Za-z_0-9]`, exp[0])
		return alpha
	}
	alpha, _ := regexp.MatchString(`[A-Za-z_0-9]`, lastChar)
	return alpha
}

type Token struct {
	Type   string
	Start  int
	End    int
	Lexeme string
}

func isNum() bool {
	num, _ := regexp.MatchString(`[0-9]`, lastChar)
	return num
}

func Lexer() Token {
	if lastChar == "\n" {
		return Token{Type: "New-Line", Start: Pointer, End: Pointer, Lexeme: "NEWLINE"}
	}
	for isSpace() || lastChar == "" {
		GetChar()
	}

	if isAlpha() {
		start := Pointer
		id := ""
		for isAlphaNum() == true {
			id += lastChar
			if isAlphaNum(Peek()) == false || Peek() == "EOF" {
				break
			}
			GetChar()
		}
		if id == "func" {
			return Token{Type: "Keyword", Start: start, End: Pointer, Lexeme: id}
		}

		if id == "string" {
			return Token{Type: "Keyword-Type", Start: start, End: Pointer, Lexeme: id}
		}

		if id == "int" {
			return Token{Type: "Keyword-Type", Start: start, End: Pointer, Lexeme: id}
		}

		if id == "float" {
			return Token{Type: "Keyword-Type", Start: start, End: Pointer, Lexeme: id}
		}

		return Token{Type: "Id", Start: start, End: Pointer, Lexeme: id}
	}
	if lastChar == "-" && Peek() == ">" {
		start := Pointer
		GetChar()
		return Token{Type: "Assign", Start: start, End: Pointer, Lexeme: "->"}
	}
	if lastChar == "(" {
		start := Pointer
		return Token{Type: "L-Bracket", Start: start, End: Pointer, Lexeme: "("}
	}
	if lastChar == ")" {
		start := Pointer
		return Token{Type: "R-Bracket", Start: start, End: Pointer, Lexeme: ")"}
	}
	if isNum() || lastChar == "." {
		start := Pointer
		num := lastChar
		GetChar()
		for {
			num += lastChar
			if Peek() == "\n" || isNum() == false || lastChar == "." {
				break
			}
			GetChar()
		}
		if num == "." {
			return Token{Type: "EndState", Start: start, End: Pointer, Lexeme: "."}
		}
		if strings.Contains(num, ".") {
			return Token{Type: "Float", Start: start, End: Pointer, Lexeme: num}
		}
		return Token{Type: "Number", Start: start, End: Pointer, Lexeme: num}
	}

	if lastChar == "\"" {
		start := Pointer
		str := ""
		for {
			if start == Pointer {
				GetChar()
			}
			if lastChar != "\"" {
				str += lastChar
				GetChar()
			} else {
				break
			}
		}
		return Token{Type: "String", Start: start, End: Pointer, Lexeme: str}
	}
	if lastChar == ":" && Peek() == ":" {
		start := Pointer
		GetChar()
		return Token{Type: "ref", Start: start, End: Pointer, Lexeme: "::"}
	}
	if lastChar == "EOF" {
		return Token{Type: "End", Start: Pointer, End: Pointer, Lexeme: "EOF"}
	}

	if lastChar == "=" && Peek() == "=" {
		GetChar()
		return Token{Type: "compare", Start: Pointer, End: Pointer, Lexeme: "=="}
	} else if lastChar == "<" && Peek() == "=" {
		GetChar()
		return Token{Type: "compare", Start: Pointer, End: Pointer, Lexeme: "<="}
	} else if lastChar == ">" && Peek() == "=" {
		GetChar()
		return Token{Type: "compare", Start: Pointer, End: Pointer, Lexeme: ">="}
	}
	return Token{Type: "Any", Start: Pointer, End: Pointer, Lexeme: lastChar}
}
