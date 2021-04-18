package parse

import (
	"kall/internal/lex"
	"kall/utils"
	"strings"
)

var Toks []lex.Token

var curTok lex.Token
var Pointer int = -1
var line int = 1

type Var struct {
	Name  string
	Type  string
	Value []Item
}

type Func struct {
	Name       string
	Type       string
	Body       []Item
	ReturnType string
	Args       []Item
}

type String struct {
	Value  string
	Length int
}
type Ops struct {
	Type  string
	Value string
}
type Equation struct {
	Body [][]Ops
}

type Item struct {
	Type     string
	Variable *Var
	String   *String
	Function *Func
	Equation *Equation
	Line     int
	IsEOF    bool
}

func getTok() {
	if Pointer == len(Toks)-1 {
		curTok = lex.Token{Type: "EOF"}
	} else {
		Pointer += 1
		curTok = Toks[Pointer]
	}
}

func peek() lex.Token {
	if Pointer+1 == len(Toks) {
		return lex.Token{Type: "EOF", Lexeme: "EOF"}
	}
	return Toks[Pointer+1]
}

func typeParser() Item {
	typename := curTok.Lexeme
	getTok()
	if curTok.Type == "Id" {
		name := curTok.Lexeme
		body := []Item{}
		getTok()
		if curTok.Type == "Assign" {
			getTok()
			for {
				tok := Parse(false)
				body = append(body, tok)
				if peek().Lexeme == "NEWLINE" {
					break
				}
				if peek().Lexeme == "EOF" {
					break
				}
				getTok()
			}
			ret := Item{Line: line}
			ret.Type = "variable"
			ret.Variable = &Var{Name: name, Value: body, Type: typename}
			return ret
		} else {
			ret := Item{Line: line}
			ret.Type = "variable-ref"
			ret.Variable = &Var{Name: name, Type: typename}
			return ret
		}
	} else {
		last := Toks[Pointer-1]
		str := strings.Join(lex.Chars[last.Start:curTok.End], "")
		problem := strings.Join(lex.Chars[curTok.Start:curTok.End], "")
		start := strings.Index(str, problem)
		end := start + len(problem)
		utils.MakeErr("Unexpected token", str, start, end, line)
	}
	return Item{}
}

func stringParser() Item {
	ret := Item{Line: line}
	ret.String = &String{Value: curTok.Lexeme, Length: len(curTok.Lexeme)}
	ret.Type = "string"
	return ret
}

func EOFParser() Item {
	ret := Item{Line: line, IsEOF: true}
	return ret
}

func funcParser() Item {
	getTok() // Eat func keyword
	name := curTok.Lexeme
	args := []Item{}
	returnType := ""
	getTok()
	if curTok.Lexeme != "(" {
		last := Toks[Pointer-1]
		str := strings.Join(lex.Chars[last.Start:curTok.End], "")
		problem := strings.Join(lex.Chars[curTok.Start:curTok.End], "")
		start := strings.Index(str, problem)
		end := start + len(problem)
		utils.MakeErr("Expected token (", str, start, end, line)
	}
	getTok() // Eat (
	for curTok.Lexeme != ")" {
		if curTok.Lexeme == "," {
			continue
		}
		args = append(args, Parse(false))
		if curTok.Lexeme == ")" {
			break
		}
		getTok()
	}
	getTok()

	if curTok.Lexeme != "->" {
		last := Toks[Pointer-1]
		str := strings.Join(lex.Chars[last.Start:curTok.End], "")
		problem := strings.Join(lex.Chars[curTok.Start:curTok.End], "")
		start := strings.Index(str, problem)
		end := start + len(problem)
		utils.MakeErr("Expected token ->", str, start, end, line)
	}
	getTok()

	if curTok.Type != "Id" && curTok.Type != "Keyword-Type" {
		last := Toks[Pointer-1]
		str := strings.Join(lex.Chars[last.Start:curTok.End], "")
		problem := strings.Join(lex.Chars[curTok.Start:curTok.End], "")
		start := strings.Index(str, problem)
		end := start + len(problem)
		utils.MakeErr("Invalid return type "+curTok.Lexeme, str, start, end, line)
	}
	returnType = curTok.Lexeme
	getTok()

	if curTok.Lexeme != "{" {
		last := Toks[Pointer-1]
		str := strings.Join(lex.Chars[last.Start:curTok.End], "")
		problem := strings.Join(lex.Chars[curTok.Start:curTok.End], "")
		start := strings.Index(str, problem)
		end := start + len(problem)
		utils.MakeErr("Expected token {", str, start, end, line)
	}

	getTok() // eat {
	body := []Item{}
	for curTok.Lexeme != "}" {
		body = append(body, Parse(false))
		if curTok.Lexeme == "}" {
			break
		}
		getTok()
	}

	ret := Item{Line: line}
	ret.Function = &Func{Name: name, Type: "function", ReturnType: returnType, Args: args, Body: body}
	ret.Type = "function"
	return ret
}

func Parse(move bool) Item {
	if move {
		getTok()
	}
	if curTok.Lexeme == "func" {
		return funcParser()
	}
	if curTok.Type == "Keyword-Type" {
		return typeParser()
	}
	if curTok.Type == "String" {
		return stringParser()
	}
	if curTok.Type == "Number" || curTok.Type == "Float" {
		return PresParse()
	}
	if curTok.Type == "New-Line" {
		line += 1
	}
	if curTok.Type == "EOF" {
		return EOFParser()
	}
	return Item{}
}
