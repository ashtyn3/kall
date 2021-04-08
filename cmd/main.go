package main

import (
	"fmt"
	"kall/internal/lex"
	"kall/internal/parse"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./hi.kll")
	arr := strings.Split(string(data), "")
	lex.Chars = arr
	tokArr := []lex.Token{}
	for lex.Peek() != "EOF" {
		tok := lex.Lexer()
		tokArr = append(tokArr, tok)
		lex.GetChar()
	}
	parse.Toks = tokArr

	tokens := []parse.Item{}
	for {
		out := parse.Parse(true)
		tokens = append(tokens, out)
		if out.IsEOF == true {
			break
		}
	}
	fmt.Println(tokens[0].Function.Body)
}
