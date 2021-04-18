package parse

import (
	"fmt"
	"kall/internal/lex"
	"regexp"
	"strings"

	"github.com/robertkrimen/otto/parser"
)

func PresParse() Item {
	u := []string{}
	fmt.Println(Toks)
	for {
		u = append(u, curTok.Lexeme)
		if peek().Type != "Number" && peek().Type != "Id" && peek().Lexeme != "+" && peek().Lexeme != "-" && peek().Lexeme != "*" && peek().Lexeme != "/" && peek().Lexeme != "(" && peek().Lexeme != ")" || peek().Lexeme == "NEWLINE" {
			break
		}
		getTok()
	}
	for _, op := range u {
		m, _ := regexp.MatchString(`[A-Za-z_][A-Za-z0-9_]+`, op)
		if m {
			op = "hi"
		}
	}
	r := parser.NewParser(strings.Join(u, " "))
	s := fmt.Sprintf("%s", r.Parse())
	s = strings.TrimPrefix(s, "(")
	s = strings.TrimSuffix(s, ")")

	//old := lex.Chars
	//oldPointer := lex.Pointer
	lex.Chars = strings.Split(s, "")
	lex.Pointer = -1
	tokArr := []lex.Token{}
	for lex.Peek() != "EOF" {
		tok := lex.Lexer()
		tokArr = append(tokArr, tok)
		lex.GetChar()
	}
	toks := [][]Ops{[]Ops{}}
	place := 0
	for _, p := range tokArr {
		if p.Type == "Number" {
			toks[place] = append(toks[place], Ops{Value: p.Lexeme, Type: "num"})
		} else if p.Lexeme == "+" {
			toks[place] = append(toks[place], Ops{Value: p.Lexeme, Type: "add"})
		} else if p.Lexeme == "-" {
			toks[place] = append(toks[place], Ops{Value: p.Lexeme, Type: "sub"})
		} else if p.Lexeme == "*" {
			toks[place] = append(toks[place], Ops{Value: p.Lexeme, Type: "mul"})
		} else if p.Lexeme == "/" {
			toks[place] = append(toks[place], Ops{Value: p.Lexeme, Type: "div"})
		} else if p.Lexeme == "^" {
			toks[place] = append(toks[place], Ops{Value: p.Lexeme, Type: "pow"})
		} else if p.Lexeme == "(" {
			toks = append(toks, []Ops{})
			place++
		} else if p.Lexeme == ")" {
			place--
		} else {
			toks[place] = append(toks[place], Ops{Value: p.Lexeme, Type: "other"})
		}
	}
	eq := &Equation{Body: toks}
	return Item{Type: "equation", Equation: eq, Line: line}
	//lex.Chars = old
	//lex.Pointer = oldPointer

}
