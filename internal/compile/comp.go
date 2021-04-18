package compile

import (
	"kall/internal/parse"
	"sort"
)

var context *Module = NewModule("main")
var Scope map[string][]parse.Item = make(map[string][]parse.Item)

func eqRes(eq *parse.Equation) string {
	r := ""
	for i := len(eq.Body) - 1; i >= 0; i-- {
		for _, p := range eq.Body[i] {
			if p.Value != "NEWLINE" {
				r += p.Value
			}
		}
	}
	return r
}

func Compiler(tree []parse.Item, scope string) {
	for _, b := range tree {
		if b.Type == "function" {
			Scope[b.Function.Name] = append(Scope[b.Function.Name], b.Function.Args...)
			Compiler(b.Function.Body, b.Function.Name)
		}
		if b.Type == "variable" {
			i := sort.Search(len(Scope[scope]), func(i int) bool { return Scope[scope][i].Variable.Name == b.Variable.Name })
			if i == len(Scope[scope]) {
				if b.Variable.Type == "float" {
					for _, v := range b.Variable.Value {
						if v.Type == "equation" {
							eqRes(b.Variable.Value[0].Equation)
						}
					}
					//context.NewFloat(b.Variable.Value[0])
					//context.NewVariable(b.Variable.Name)
				}
			}
		}
	}
}
