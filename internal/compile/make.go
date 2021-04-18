package compile

import "fmt"

type FloatType struct {
	Size float64
}

type IntType struct {
	Size int64
}
type CharsType struct {
	Size  float64
	chars []byte
}

type Types struct {
	Name   string
	Int    IntType
	Float  FloatType
	String CharsType
}

type VariableTemp struct {
	Name  string
	Value *Types
	Code  string
}

type Params struct {
	Bind bool
}

type Module struct {
	Name    string
	imports []string
	Body    []string
}

func (m *Module) NewFloat(val float64) *Types {
	return &Types{Name: "f32", Float: FloatType{Size: val}}
}
func (m *Module) NewInt(val int64) *Types {
	return &Types{Name: "u32", Int: IntType{Size: val}}
}
func (m *Module) NewVariable(name string, varType *Types, options ...Params) *VariableTemp {
	var code string
	if varType.Name == "string" {
		varType.Name = "str"
	}
	if len(options) > 0 {
		if options[0].Bind == false {
			code = "let " + name + ":" + varType.Name
		} else {
			code = "let mut " + name + ":" + varType.Name
		}
	} else {
		code = "let " + name + ":" + varType.Name
	}
	if varType.Name == "f32" {
		f := float64(varType.Float.Size)
		str := fmt.Sprintf("%f", f)
		code += " = " + str + ";"
		m.Body = append(m.Body, code)
		return &VariableTemp{Name: name, Value: varType, Code: code}
	}

	if varType.Name == "string" || varType.Name == "byte" {
		if varType.Name == "string" {
			code += " = \"" + string(varType.String.chars) + "\";"
		}
		return &VariableTemp{Name: name, Value: varType, Code: code}
	}
	return &VariableTemp{}
}
func NewModule(name string) *Module {
	return &Module{Name: name, Body: []string{"fn main() {"}}
}

func (m *Module) Compile() {
	m.Body = append(m.Body, "}")
}
